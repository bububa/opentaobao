package legalsuit

// SealTaskModel 结构体
type SealTaskModel struct {
	// 参数对象
	SealFileModels []SealFileModel `json:"seal_file_models,omitempty" xml:"seal_file_models>seal_file_model,omitempty"`
	// 推送人
	PushPeople string `json:"push_people,omitempty" xml:"push_people,omitempty"`
	// 用印类型
	SealType string `json:"seal_type,omitempty" xml:"seal_type,omitempty"`
	// 处理人
	HandlerWorkNo string `json:"handler_work_no,omitempty" xml:"handler_work_no,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
}
