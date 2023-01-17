package aedropshiper

// AeMultimediaInfoDto 结构体
type AeMultimediaInfoDto struct {
	// Video information
	AeVideoDtos []AeVideoDto `json:"ae_video_dtos,omitempty" xml:"ae_video_dtos>ae_video_dto,omitempty"`
	// List of main images of the product
	ImageUrls string `json:"image_urls,omitempty" xml:"image_urls,omitempty"`
}
