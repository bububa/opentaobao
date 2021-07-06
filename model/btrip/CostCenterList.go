package btrip

// CostCenterList 结构体
type CostCenterList struct {
	// corpId
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// number
	Number string `json:"number,omitempty" xml:"number,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
