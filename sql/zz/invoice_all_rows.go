// Code generated by "sqlc-gen-zombiezen". DO NOT EDIT.

package zz

import (
	"fmt"

	"zombiezen.com/go/sqlite"

	"time"

	"github.com/delaneyj/toolbelt"
)

type InvoiceAllRowsRes struct {
	Id        int64     `json:"id"`
	InvoiceId int64     `json:"invoice_id"`
	Number    int64     `json:"number"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Comment   string    `json:"comment"`
	Uid       string    `json:"uid"`
	Qty       float64   `json:"qty"`
	Units     string    `json:"units"`
	Taxid     float64   `json:"taxid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InvoiceAllRowsStmt struct {
	stmt *sqlite.Stmt
}

func InvoiceAllRows(tx *sqlite.Conn) *InvoiceAllRowsStmt {
	// Prepare the statement into connection cache
	stmt := tx.Prep(`
select id, invoice_id, number, name, price, comment, uid, qty, units, taxid, created_at, updated_at from invoice_rows where invoice_id = ?1
    `)
	ps := &InvoiceAllRowsStmt{stmt: stmt}
	return ps
}

func (ps *InvoiceAllRowsStmt) Run(
	id int64,
) (
	res []InvoiceAllRowsRes,
	err error,
) {
	defer ps.stmt.Reset()

	// Bind parameters
	ps.stmt.BindInt64(1, id)

	// Execute the query
	for {
		if hasRow, err := ps.stmt.Step(); err != nil {
			return res, fmt.Errorf("failed to execute {{.Name.Lower}} SQL: %w", err)
		} else if !hasRow {
			break
		}

		row := InvoiceAllRowsRes{}
		row.Id = ps.stmt.ColumnInt64(0)
		row.InvoiceId = ps.stmt.ColumnInt64(1)
		row.Number = ps.stmt.ColumnInt64(2)
		row.Name = ps.stmt.ColumnText(3)
		row.Price = ps.stmt.ColumnFloat(4)
		row.Comment = ps.stmt.ColumnText(5)
		row.Uid = ps.stmt.ColumnText(6)
		row.Qty = ps.stmt.ColumnFloat(7)
		row.Units = ps.stmt.ColumnText(8)
		row.Taxid = ps.stmt.ColumnFloat(9)
		row.CreatedAt = toolbelt.JulianDayToTime(ps.stmt.ColumnFloat(10))
		row.UpdatedAt = toolbelt.JulianDayToTime(ps.stmt.ColumnFloat(11))
		res = append(res, row)
	}

	return res, nil
}

func OnceInvoiceAllRows(
	tx *sqlite.Conn,
	id int64,
) (
	res []InvoiceAllRowsRes,
	err error,
) {
	ps := InvoiceAllRows(tx)

	return ps.Run(
		id,
	)
}