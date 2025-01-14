/* Package repositories provides the implementation of the repositories interfaces. */ package repositories

import (
	"database/sql"
	"fmt"
	"rest_api_go/models"
	"strings"
)

type PostgresDeviceRepository struct {
	DB *sql.DB
}

func NewPostgresDeviceRepository(db *sql.DB) *PostgresDeviceRepository {
	return &PostgresDeviceRepository{DB: db}
}

// GetAllDevices returns all devices from the database.
func (repo *PostgresDeviceRepository) GetAllDevices() ([]models.Device, error) {
	query := `SELECT id, name, version FROM devices`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var dev models.Device
		if err := rows.Scan(&dev.ID, &dev.Name, &dev.Version); err != nil {
			return nil, err
		}
		devices = append(devices, dev)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

// GetDeviceByID returns a device with the given ID.
func (repo *PostgresDeviceRepository) GetDeviceByID(id int) (*models.Device, error) {
	query := `SELECT id, name, version FROM devices WHERE id = $1`
	var dev models.Device
	err := repo.DB.QueryRow(query, id).Scan(&dev.ID, &dev.Name, &dev.Version)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("device with ID %d not found", id)
		}
		return nil, err
	}

	return &dev, nil
}

// DeleteDevice deletes a device with the given ID.
func (repo *PostgresDeviceRepository) UpdateDevice(dev models.Device) error {
	query := `UPDATE devices SET name = $1, version = $2 WHERE id = $3`
	_, err := repo.DB.Exec(query, dev.Name, dev.Version, dev.ID)
	if err != nil {
		return err
	}
	return nil
}

// CreateDevice creates a new device in the database.
func (repo *PostgresDeviceRepository) CreateDevice(dev models.Device) (int, error) {
	query := `INSERT INTO devices (name, version) VALUES ($1, $2) RETURNING id`
	var id int
	err := repo.DB.QueryRow(query, dev.Name, dev.Version).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *PostgresDeviceRepository) PatchDevice(id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	fieldsToUpdate := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1) // +1 for the ID
	index := 1

	for field, value := range updates {
		fieldsToUpdate = append(fieldsToUpdate, fmt.Sprintf("%s = $%d", field, index))
		args = append(args, value)
		index++
	}

	// Add the ID to the end for the WHERE clause
	args = append(args, id)

	query := fmt.Sprintf(`UPDATE devices SET %s WHERE id = $%d`,
		strings.Join(fieldsToUpdate, ", "), index)

	// Execute the composed query with arguments
	_, err := repo.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("unable to update device: %w", err)
	}
	return nil
}

func (repo *PostgresDeviceRepository) DeleteDevice(id int) error {
	query := `DELETE FROM devices WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
