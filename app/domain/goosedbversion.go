// Package domain contains the types for schema 'public'.
package domain

import "github.com/lib/pq"

// Code generated by xo. DO NOT EDIT.

// GooseDbVersion represents a row from 'public.goose_db_version'.
type GooseDbVersion struct {
	ID        int         `json:"id"`         // id
	VersionID int64       `json:"version_id"` // version_id
	IsApplied bool        `json:"is_applied"` // is_applied
	Tstamp    pq.NullTime `json:"tstamp"`     // tstamp

}

type GooseDbVersionService interface {
	DoesGooseDbVersionExists(gdv *GooseDbVersion) (bool, error)
	InsertGooseDbVersion(gdv *GooseDbVersion) error
	UpdateGooseDbVersion(gdv *GooseDbVersion) error
	UpsertGooseDbVersion(gdv *GooseDbVersion) error
	DeleteGooseDbVersion(gdv *GooseDbVersion) error
	GetAllGooseDbVersions() ([]*GooseDbVersion, error)
	GetChunkedGooseDbVersions(limit int, offset int) ([]*GooseDbVersion, error)
}

type GooseDbVersionServiceImpl struct {
	DB XODB
}

// Exists determines if the GooseDbVersion exists in the database.
func (serviceImpl *GooseDbVersionServiceImpl) DoesGooseDbVersionExists(gdv *GooseDbVersion) (bool, error) {
	panic("not yet implemented")
}

// Insert inserts the GooseDbVersion to the database.
func (serviceImpl *GooseDbVersionServiceImpl) InsertGooseDbVersion(gdv *GooseDbVersion) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.goose_db_version (` +
		`version_id, is_applied, tstamp` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	err = serviceImpl.DB.QueryRow(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp).Scan(&gdv.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the GooseDbVersion in the database.
func (serviceImpl *GooseDbVersionServiceImpl) UpdateGooseDbVersion(gdv *GooseDbVersion) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.goose_db_version SET (` +
		`version_id, is_applied, tstamp` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp, gdv.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp, gdv.ID)
	return err
}

// Save saves the GooseDbVersion to the database.
/*
	func (gdv *GooseDbVersion) Save(db XODB) error {
		if gdv.Exists() {
			return gdv.Update(db)
		}

		return gdv.Insert(db)
	}
*/

// Upsert performs an upsert for GooseDbVersion.
//
// NOTE: PostgreSQL 9.5+ only
func (serviceImpl *GooseDbVersionServiceImpl) UpsertGooseDbVersion(gdv *GooseDbVersion) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.goose_db_version (` +
		`id, version_id, is_applied, tstamp` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, version_id, is_applied, tstamp` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.version_id, EXCLUDED.is_applied, EXCLUDED.tstamp` +
		`)`

	// run query
	XOLog(sqlstr, gdv.ID, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	_, err = serviceImpl.DB.Exec(sqlstr, gdv.ID, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the GooseDbVersion from the database.
func (serviceImpl *GooseDbVersionServiceImpl) DeleteGooseDbVersion(gdv *GooseDbVersion) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.goose_db_version WHERE id = $1`

	// run query
	XOLog(sqlstr, gdv.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, gdv.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllGooseDbVersions returns all rows from 'public.goose_db_version',
// ordered by "created_at" in descending order.
func (serviceImpl *GooseDbVersionServiceImpl) GetAllGooseDbVersions() ([]*GooseDbVersion, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.goose_db_version`

	q, err := serviceImpl.DB.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*GooseDbVersion
	for q.Next() {
		gdv := GooseDbVersion{}

		// scan
		err = q.Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
		if err != nil {
			return nil, err
		}

		res = append(res, &gdv)
	}

	return res, nil
}

// GetChunkedGooseDbVersions returns pagingated rows from 'public.goose_db_version',
// ordered by "created_at" in descending order.
func (serviceImpl *GooseDbVersionServiceImpl) GetChunkedGooseDbVersions(limit int, offset int) ([]*GooseDbVersion, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.goose_db_version LIMIT $1 OFFSET $2`

	q, err := serviceImpl.DB.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*GooseDbVersion
	for q.Next() {
		gdv := GooseDbVersion{}

		// scan
		err = q.Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
		if err != nil {
			return nil, err
		}

		res = append(res, &gdv)
	}

	return res, nil
}

// GooseDbVersionByID retrieves a row from 'public.goose_db_version' as a GooseDbVersion.
//
// Generated from index 'goose_db_version_pkey'.
func (serviceImpl *GooseDbVersionServiceImpl) GooseDbVersionByID(_, id int) (*GooseDbVersion, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, version_id, is_applied, tstamp ` +
		`FROM public.goose_db_version ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)

	gdv := GooseDbVersion{}

	err = serviceImpl.DB.QueryRow(sqlstr, id).Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
	if err != nil {
		return nil, err
	}

	return &gdv, nil
}
