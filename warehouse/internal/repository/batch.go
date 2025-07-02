package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetAllBatches() (batches []models.Batch, err error) {

	err = db.GetDBConn().Select(&batches, `SELECT b.date, b.type,
						cp.name AS counterparty_name, cp.contact, cp.phone, cp.email,
						p.article, p.name AS product_name, p.price,
						s.quantity,
						c.zone, c.row, c.adress_code,
						u.username, u.role, u.full_name, u.active
					FROM batches b
					INNER JOIN counterparties cp ON b.counterparty_name = cp.name
					INNER JOIN products p ON b.article = p.article
					INNER JOIN storages s ON b.adress_code = s.adress_code
					INNER JOIN cells c ON  b.adress_code = c.adress_code
					INNER JOIN users u ON b.username = u.username;`)

	if err != nil {
		logger.Error.
			Printf("[repository] GetAllBatches(): error duriing getting from database: %s\n", err.Error())
		return []models.Batch{}, translateError(err)
	}

	return batches, nil
}

func GetBatchByID(id int) (batch models.Batch, err error) {

	err = db.GetDBConn().Get(&batch, `SELECT b.id, b.date, b.type,
						cp.name AS counterparty_name, cp.contact, cp.phone, cp.email,
						p.article, p.name AS product_name, p.price,
						s.quantity,
						c.zone, c.row, c.adress_code,
						u.username, u.role, u.full_name, u.active
					FROM batches b
					INNER JOIN counterparties cp ON b.counterparty_name = cp.name
					INNER JOIN products p ON b.article = p.article
					INNER JOIN storages s ON b.adress_code = s.adress_code
					INNER JOIN cells c ON  b.adress_code = c.adress_code
					INNER JOIN users u ON b.username = u.username
					WHERE b.id = $1;`, id)

	if err != nil {
		logger.Error.
			Printf("[repository] GetBatchByID(): error duriing getting from database: %s\n", err.Error())
		return models.Batch{}, translateError(err)
	}

	return batch, nil
}

func CreateBatch(b models.Batch) error {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO batches(type, counterparty_name, article, quantity, adress_code, username)
			values($1, $2, $3, $4, $5, $6);`, b.Type, b.CounterpartyName, b.Article, b.Quantity, b.AdressCode, b.Username)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateBatch(): error duriing creating batch from database: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}
