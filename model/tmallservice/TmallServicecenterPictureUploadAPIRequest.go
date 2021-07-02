package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterPictureUploadAPIRequest 上传图片 API请求
// tmall.servicecenter.picture.upload
//
// 给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
type TmallServicecenterPictureUploadAPIRequest struct {
	model.Params
	// 图片文件二进制流
	_img *model.File
	// 图片全称包括扩展名。目前支持 jpg jpeg png
	_pictureName string
	// true返回Https地址
	_isHttps bool
}

// NewTmallServicecenterPictureUploadRequest 初始化TmallServicecenterPictureUploadAPIRequest对象
func NewTmallServicecenterPictureUploadRequest() *TmallServicecenterPictureUploadAPIRequest {
	return &TmallServicecenterPictureUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterPictureUploadAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterPictureUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Img Setter
// 图片文件二进制流
func (r *TmallServicecenterPictureUploadAPIRequest) SetImg(_img *model.File) error {
	r._img = _img
	r.Set("img", _img)
	return nil
}

// Get Img Getter
func (r TmallServicecenterPictureUploadAPIRequest) GetImg() *model.File {
	return r._img
}

// Set is PictureName Setter
// 图片全称包括扩展名。目前支持 jpg jpeg png
func (r *TmallServicecenterPictureUploadAPIRequest) SetPictureName(_pictureName string) error {
	r._pictureName = _pictureName
	r.Set("picture_name", _pictureName)
	return nil
}

// Get PictureName Getter
func (r TmallServicecenterPictureUploadAPIRequest) GetPictureName() string {
	return r._pictureName
}

// Set is IsHttps Setter
// true返回Https地址
func (r *TmallServicecenterPictureUploadAPIRequest) SetIsHttps(_isHttps bool) error {
	r._isHttps = _isHttps
	r.Set("is_https", _isHttps)
	return nil
}

// Get IsHttps Getter
func (r TmallServicecenterPictureUploadAPIRequest) GetIsHttps() bool {
	return r._isHttps
}
