package legalsuit

// SealFileModel 结构体
type SealFileModel struct {
	// 文件对象
	FileModels []FileModel `json:"file_models,omitempty" xml:"file_models>file_model,omitempty"`
	// 申请说明
	Explain string `json:"explain,omitempty" xml:"explain,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 文件类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}
