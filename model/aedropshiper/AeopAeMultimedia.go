package aedropshiper

// AeopAeMultimedia 结构体
type AeopAeMultimedia struct {
	// 多媒体信息。
	AeopAEVideos []AeopAeVideo `json:"aeop_a_e_videos,omitempty" xml:"aeop_a_e_videos>aeop_ae_video,omitempty"`
}
