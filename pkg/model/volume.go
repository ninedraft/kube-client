package model

// Volume -- volume representation
//
//swagger:model
type Volume struct {
	ID          string                     `json:"id,omitempty"`
	OwnerLogin  string                     `json:"owner_login,omitempty"`
	Label       string                     `json:"label,omitempty"`
	Access      string                     `json:"access,omitempty"`
	Capacity    uint                       `json:"capacity,omitempty"`
	StorageName string                     `json:"storage_name,omitempty"` //AKA StorageClass
	AccessMode  PersistentVolumeAccessMode `json:"access_mode,omitempty"`
	CreatedAt   *string                    `json:"created_at,omitempty"`
	Owner       string                     `json:"owner,omitempty"`
}

// CreateVolume --
//swagger:ignore
type CreateVolume struct {
	TariffID string `json:"tariff-id"`
	Label    string `json:"label"`
}

// ResourceUpdateName -- contains new resource name
//swagger:ignore
type ResourceUpdateName struct {
	Label string `json:"label"`
}

type PersistentVolumeAccessMode string

const (
	// can be mounted read/write mode to exactly 1 host
	ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
	// can be mounted in read-only mode to many hosts
	ReadOnlyMany PersistentVolumeAccessMode = "ReadOnlyMany"
	// can be mounted in read/write mode to many hosts
	ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
)

// Mask removes information not interesting for users
func (vol *Volume) Mask() {
	vol.Owner = ""
}
