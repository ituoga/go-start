// Code generated by "sqlc-gen-zombiezen". DO NOT EDIT.

package zz

import (
	"fmt"
	"time"

	"github.com/delaneyj/toolbelt"
	"zombiezen.com/go/sqlite"
)

type ProductModel struct {
	Id        int64     `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductStmt struct {
	stmt *sqlite.Stmt
}

func CreateProduct(tx *sqlite.Conn) *CreateProductStmt {
	stmt := tx.Prep(`
INSERT INTO products (
        id,
        uuid,
        name,
        code,
        created_at,
        updated_at
) VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
)
    `)
	return &CreateProductStmt{stmt: stmt}
}

func (ps *CreateProductStmt) Run(m *ProductModel) error {
	defer ps.stmt.Reset()

	// Bind parameters
	ps.stmt.BindInt64(1, m.Id)

	ps.stmt.BindText(2, m.Uuid)

	ps.stmt.BindText(3, m.Name)

	ps.stmt.BindText(4, m.Code)

	ps.stmt.BindFloat(5, toolbelt.TimeToJulianDay(m.CreatedAt))

	ps.stmt.BindFloat(6, toolbelt.TimeToJulianDay(m.UpdatedAt))

	if _, err := ps.stmt.Step(); err != nil {
		return fmt.Errorf("failed to insert products: %w", err)
	}

	return nil
}

func OnceCreateProduct(tx *sqlite.Conn, m *ProductModel) error {
	ps := CreateProduct(tx)
	return ps.Run(m)
}

type ReadAllProductsStmt struct {
	stmt *sqlite.Stmt
}

func ReadAllProducts(tx *sqlite.Conn) *ReadAllProductsStmt {
	stmt := tx.Prep(`
SELECT
        id,
        uuid,
        name,
        code,
        created_at,
        updated_at
FROM products
    `)
	return &ReadAllProductsStmt{stmt: stmt}
}

func (ps *ReadAllProductsStmt) Run() ([]*ProductModel, error) {
	defer ps.stmt.Reset()

	var models []*ProductModel
	for {
		hasRow, err := ps.stmt.Step()
		if err != nil {
			return nil, fmt.Errorf("failed to read products: %w", err)
		} else if !hasRow {
			break
		}

		m := &ProductModel{}

		m.Id = ps.stmt.ColumnInt64(0)

		m.Uuid = ps.stmt.ColumnText(1)

		m.Name = ps.stmt.ColumnText(2)

		m.Code = ps.stmt.ColumnText(3)

		m.CreatedAt = toolbelt.JulianDayToTime(ps.stmt.ColumnFloat(4))

		m.UpdatedAt = toolbelt.JulianDayToTime(ps.stmt.ColumnFloat(5))

		models = append(models, m)
	}

	return models, nil
}

func OnceReadAllProducts(tx *sqlite.Conn) ([]*ProductModel, error) {
	ps := ReadAllProducts(tx)
	return ps.Run()
}

type ReadByIDProductStmt struct {
	stmt *sqlite.Stmt
}

func ReadByIDProduct(tx *sqlite.Conn) *ReadByIDProductStmt {
	stmt := tx.Prep(`
SELECT
        id,
        uuid,
        name,
        code,
        created_at,
        updated_at
FROM products
WHERE id = ?
    `)
	return &ReadByIDProductStmt{stmt: stmt}
}

func (ps *ReadByIDProductStmt) Run(id int64) (*ProductModel, error) {
	defer ps.stmt.Reset()

	ps.stmt.BindInt64(1, id)

	if hasRow, err := ps.stmt.Step(); err != nil {
		return nil, fmt.Errorf("failed to read products: %w", err)
	} else if !hasRow {
		return nil, nil
	}

	m := &ProductModel{}

	m.Id = ps.stmt.ColumnInt64(0)

	m.Uuid = ps.stmt.ColumnText(1)

	m.Name = ps.stmt.ColumnText(2)

	m.Code = ps.stmt.ColumnText(3)

	m.CreatedAt = toolbelt.JulianDayToTime(ps.stmt.ColumnFloat(4))

	m.UpdatedAt = toolbelt.JulianDayToTime(ps.stmt.ColumnFloat(5))

	return m, nil
}

func OnceReadByIDProduct(tx *sqlite.Conn, id int64) (*ProductModel, error) {
	ps := ReadByIDProduct(tx)
	return ps.Run(id)
}

func CountProducts(tx *sqlite.Conn) (int64, error) {
	stmt := tx.Prep(`
SELECT COUNT(*)
FROM products
    `)
	defer stmt.Reset()

	if hasRow, err := stmt.Step(); err != nil {
		return 0, fmt.Errorf("failed to count products: %w", err)
	} else if !hasRow {
		return 0, nil
	}

	return stmt.ColumnInt64(0), nil
}

func OnceCountProducts(tx *sqlite.Conn) (int64, error) {
	return CountProducts(tx)
}

type UpdateProductStmt struct {
	stmt *sqlite.Stmt
}

func UpdateProduct(tx *sqlite.Conn) *UpdateProductStmt {
	stmt := tx.Prep(`
UPDATE products
SET
        uuid = ?2,
        name = ?3,
        code = ?4,
        created_at = ?5,
        updated_at = ?6
WHERE id = ?1
    `)
	return &UpdateProductStmt{stmt: stmt}
}

func (ps *UpdateProductStmt) Run(m *ProductModel) error {
	defer ps.stmt.Reset()

	// Bind parameters
	ps.stmt.BindInt64(1, m.Id)

	ps.stmt.BindText(2, m.Uuid)

	ps.stmt.BindText(3, m.Name)

	ps.stmt.BindText(4, m.Code)

	ps.stmt.BindFloat(5, toolbelt.TimeToJulianDay(m.CreatedAt))

	ps.stmt.BindFloat(6, toolbelt.TimeToJulianDay(m.UpdatedAt))

	if _, err := ps.stmt.Step(); err != nil {
		return fmt.Errorf("failed to update products: %w", err)
	}

	return nil
}

func OnceUpdateProduct(tx *sqlite.Conn, m *ProductModel) error {
	ps := UpdateProduct(tx)
	return ps.Run(m)
}

type DeleteProductStmt struct {
	stmt *sqlite.Stmt
}

func DeleteProduct(tx *sqlite.Conn) *DeleteProductStmt {
	stmt := tx.Prep(`
DELETE FROM products
WHERE id = ?
    `)
	return &DeleteProductStmt{stmt: stmt}
}

func (ps *DeleteProductStmt) Run(id int64) error {
	defer ps.stmt.Reset()

	ps.stmt.BindInt64(1, id)

	if _, err := ps.stmt.Step(); err != nil {
		return fmt.Errorf("failed to delete products: %w", err)
	}

	return nil
}

func OnceDeleteProduct(tx *sqlite.Conn, id int64) error {
	ps := DeleteProduct(tx)
	return ps.Run(id)
}
