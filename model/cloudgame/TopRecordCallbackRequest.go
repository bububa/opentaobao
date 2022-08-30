package cloudgame

// TopRecordCallbackRequest 结构体
type TopRecordCallbackRequest struct {
	// 图片oss地址列表
	Images []OssDto `json:"images,omitempty" xml:"images>oss_dto,omitempty"`
	// 视频oss地址列表
	Videos []OssDto `json:"videos,omitempty" xml:"videos>oss_dto,omitempty"`
	// 云游戏mixUserId
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// api版本version
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 扩展json字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}
