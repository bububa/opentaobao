package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiFacebodyDetectfaceAPIRequest 人脸检测定位 API请求
// aliyun.viapi.facebody.detectface
//
// 输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征、(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiFacebodyDetectfaceAPIRequest struct {
	model.Params
	// 图片url地址(http/https)
	_imageUrl string
	// 0: 通过url识别，参数image_url不为空；1: 通过图片content识别，参数content不为空 支持图片格式：JPEG、JPG、BMP、PNG
	_imageType int64
}

// NewAliyunViapiFacebodyDetectfaceRequest 初始化AliyunViapiFacebodyDetectfaceAPIRequest对象
func NewAliyunViapiFacebodyDetectfaceRequest() *AliyunViapiFacebodyDetectfaceAPIRequest {
	return &AliyunViapiFacebodyDetectfaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiFacebodyDetectfaceAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.facebody.detectface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiFacebodyDetectfaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiFacebodyDetectfaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 图片url地址(http/https)
func (r *AliyunViapiFacebodyDetectfaceAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiFacebodyDetectfaceAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

// SetImageType is ImageType Setter
// 0: 通过url识别，参数image_url不为空；1: 通过图片content识别，参数content不为空 支持图片格式：JPEG、JPG、BMP、PNG
func (r *AliyunViapiFacebodyDetectfaceAPIRequest) SetImageType(_imageType int64) error {
	r._imageType = _imageType
	r.Set("image_type", _imageType)
	return nil
}

// GetImageType ImageType Getter
func (r AliyunViapiFacebodyDetectfaceAPIRequest) GetImageType() int64 {
	return r._imageType
}
