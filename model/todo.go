package model

type TodoModel struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name     string `json:"name"`
	NodeID   string `json:"nodeID"`
	RemoteIP string `json:"remoteIP"`
	DateField
}
