package product

import (
	"database/sql"
	"errors"

	"github.com/imbafff/product-warehouse-api/internal/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(p *entity.Product) (int64, error) {
	query := `
		INSERT INTO products (name, description, price, quantity)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(
		query,
		p.Name,
		p.Description,
		p.Price,
		p.Quantity,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresRepository) GetByID(id int64) (*entity.Product, error) {
	query := `
		SELECT id, name, description, price, quantity
		FROM products
		WHERE id = $1
	`

	var p entity.Product

	err := r.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Quantity,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostgresRepository) Update(id int64, p *entity.Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, quantity = $4
		WHERE id = $5
	`

	res, err := r.db.Exec(
		query,
		p.Name,
		p.Description,
		p.Price,
		p.Quantity,
		id,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *PostgresRepository) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = $1`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *PostgresRepository) GetAll() ([]*entity.Product, error) {
	query := `
		SELECT id, name, description, price, quantity
		FROM products
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Quantity,
		); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}
