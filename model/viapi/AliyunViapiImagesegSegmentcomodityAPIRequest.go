package viapi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmentcomodityAPIRequest 商品分割 API请求
// aliyun.viapi.imageseg.segmentcomodity
//
// 识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmentcomodityAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiImagesegSegmentcomodityRequest 初始化AliyunViapiImagesegSegmentcomodityAPIRequest对象
func NewAliyunViapiImagesegSegmentcomodityRequest() *AliyunViapiImagesegSegmentcomodityAPIRequest {
	return &AliyunViapiImagesegSegmentcomodityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunViapiImagesegSegmentcomodityAPIRequest) Reset() {
	r._imageUrl = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segmentcomodity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmentcomodityAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

var poolAliyunViapiImagesegSegmentcomodityAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunViapiImagesegSegmentcomodityRequest()
	},
}

// GetAliyunViapiImagesegSegmentcomodityRequest 从 sync.Pool 获取 AliyunViapiImagesegSegmentcomodityAPIRequest
func GetAliyunViapiImagesegSegmentcomodityAPIRequest() *AliyunViapiImagesegSegmentcomodityAPIRequest {
	return poolAliyunViapiImagesegSegmentcomodityAPIRequest.Get().(*AliyunViapiImagesegSegmentcomodityAPIRequest)
}

// ReleaseAliyunViapiImagesegSegmentcomodityAPIRequest 将 AliyunViapiImagesegSegmentcomodityAPIRequest 放入 sync.Pool
func ReleaseAliyunViapiImagesegSegmentcomodityAPIRequest(v *AliyunViapiImagesegSegmentcomodityAPIRequest) {
	v.Reset()
	poolAliyunViapiImagesegSegmentcomodityAPIRequest.Put(v)
}
