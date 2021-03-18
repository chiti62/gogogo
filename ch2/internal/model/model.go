package model

type Model struct {
	ID          uint32 `gorm:"primary_key" json:"id"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	CreateTime  uint32 `json:"create_time"`
	UpdateTime  uint32 `json:"update_time"`
	DeletedTime uint32 `json:"delete_time"`
	IsDel       uint8  `json:"is_del"`
}
