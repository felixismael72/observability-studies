package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

/* domain */

type Author struct {
	id        uuid.UUID
	name      string
	specialty string
	createdAt time.Time
	updatedAt time.Time
}

func NewAuthor(id uuid.UUID, name, specialty string,
	createdAt, updatedAt time.Time) *Author {
	return &Author{
		id:        id,
		name:      name,
		specialty: specialty,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

/* domain */

/* repository */

type AuthorLoader interface {
	Insert(Author) (*uuid.UUID, *LibraryError)
}

var _ AuthorLoader = AuthorRepository{}

type AuthorRepository struct{}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{}
}

func (r AuthorRepository) Insert(author Author) (*uuid.UUID, *LibraryError) {
	conn, err := GetConnection()
	if err != nil {
		log.Error(err)
		return nil, NewLibraryError(503, UnavailableDataBase)
	}
	defer CloseConnection(conn)

	query := `
		insert into author(id, name, specialty, created_at, updated_at)
		values (default, $1, $2, default, default) returning id;
	`

	var authorID uuid.UUID

	err = conn.Get(&authorID, query, author.name, author.specialty)
	if err != nil {
		log.Error(err)
		return nil, NewLibraryError(500, DataOperationError)
	}

	return &authorID, nil
}

/* repository */

/* service */

type AuthorManager interface {
	Create(author Author) (*uuid.UUID, *LibraryError)
}

var _ AuthorManager = (*AuthorService)(nil)

type AuthorService struct {
	authorRepo AuthorLoader
}

func NewAuthorService(authorRepo AuthorLoader) *AuthorService {
	return &AuthorService{authorRepo: authorRepo}
}

func (s AuthorService) Create(author Author) (*uuid.UUID, *LibraryError) {
	authorID, err := s.authorRepo.Insert(author)
	if err != nil {
		return nil, err
	}

	return authorID, nil
}

/* service */

/* request */

type AuthorRequest struct {
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

func (req AuthorRequest) AsDomain() *Author {
	return &Author{
		name:      req.Name,
		specialty: req.Specialty,
	}
}

/* request */

/* handler */

type AuthorHandler struct {
	authorService AuthorManager
}

func NewAuthorHandler(authorService AuthorManager) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (h AuthorHandler) Post(context echo.Context) error {
	var (
		dto AuthorRequest
		err error
	)

	if err = context.Bind(&dto); err != nil {
		log.Error(err)
		return context.JSON(http.StatusUnprocessableEntity, UnprocessableEntity)
	}

	authorID, createErr := h.authorService.Create(*dto.AsDomain())
	if createErr != nil {
		return context.JSON(createErr.statusCode, createErr.message)
	}

	return context.JSON(http.StatusCreated, authorID)
}

/* handler */
