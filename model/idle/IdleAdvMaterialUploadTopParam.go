package idle

// IdleAdvMaterialUploadTopParam 结构体
type IdleAdvMaterialUploadTopParam struct {
	// 素材上传详细信息参数
	UploadDetails []IdleAdvMaterialUploadDetailTopParam `json:"upload_details,omitempty" xml:"upload_details>idle_adv_material_upload_detail_top_param,omitempty"`
}
