package repositories

import "rest_api_go/models"

// DeviceRepository is an interface that defines the methods that a device repository should implement.
type DeviceRepository interface {
	GetAllDevices() ([]models.Device, error)
	GetDeviceByID(id int) (*models.Device, error)
	CreateDevice(dev models.Device) (int, error)
	UpdateDevice(dev models.Device) error
	PatchDevice(id int, updates map[string]interface{}) error
	DeleteDevice(id int) error
}
