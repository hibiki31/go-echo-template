package model

type DateField struct {
	Updated int64 `json:"updated" gorm:"autoUpdateTime:milli"` // 更新時間としてUNIX msを使用する
	Created int64 `json:"created" gorm:"autoCreateTime:milli"` // 作成時間としてUNIX msを使用する
}
