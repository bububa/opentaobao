package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiFacebodyRecognizefaceAPIRequest
人脸属性识别 API请求
aliyun.viapi.facebody.recognizeface

输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiFacebodyRecognizefaceAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
	// 图片类型, ,取值范围[0:ImageURL, 1:ImageContent]
	_imageType int64
}

// NewAliyunViapiFacebodyRecognizefaceRequest 初始化AliyunViapiFacebodyRecognizefaceAPIRequest对象
func NewAliyunViapiFacebodyRecognizefaceRequest() *AliyunViapiFacebodyRecognizefaceAPIRequest {
	return &AliyunViapiFacebodyRecognizefaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiFacebodyRecognizefaceAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.facebody.recognizeface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiFacebodyRecognizefaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiFacebodyRecognizefaceAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// Get ImageUrl Getter
func (r AliyunViapiFacebodyRecognizefaceAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

// Set is ImageType Setter
// 图片类型, ,取值范围[0:ImageURL, 1:ImageContent]
func (r *AliyunViapiFacebodyRecognizefaceAPIRequest) SetImageType(_imageType int64) error {
	r._imageType = _imageType
	r.Set("image_type", _imageType)
	return nil
}

// Get ImageType Getter
func (r AliyunViapiFacebodyRecognizefaceAPIRequest) GetImageType() int64 {
	return r._imageType
}
