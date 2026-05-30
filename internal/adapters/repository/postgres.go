package repository

import (
	"context"
	"errors"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

var _ ports.LinkRepository = (*PostgresRepository)(nil)

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (p *PostgresRepository) GetByCode(ctx context.Context, code string) (domain.Link, error) {
	query := `select id, original_url, short_code, created_at, visits from links where short_code = $1`

	var link domain.Link
	err := p.pool.QueryRow(ctx, query, code).Scan(
		&link.Id,
		&link.OriginalURL,
		&link.ShortCode,
		&link.CreatedAt,
		&link.Visits,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Link{}, errors.New("not found")
		}
		return domain.Link{}, err
	}

	return link, nil
}

func (p *PostgresRepository) Save(ctx context.Context, link domain.Link) error {
	query := `
		insert into links (id, original_url, short_code, created_at, visits)
		values ($1, $2, $3, $4, $5)
	`

	_, err := p.pool.Exec(ctx, query,
		link.Id,
		link.OriginalURL,
		link.ShortCode,
		link.CreatedAt,
		link.Visits,
	)

	return err
}
