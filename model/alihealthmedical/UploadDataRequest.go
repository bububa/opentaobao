package alihealthmedical

// UploadDataRequest 结构体
type UploadDataRequest struct {
	// 文件类型名
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 医生外部id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
}
