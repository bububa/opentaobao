package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterPictureUploadAPIRequest 上传图片 API请求
// tmall.servicecenter.picture.upload
//
// 给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
type TmallServicecenterPictureUploadAPIRequest struct {
	model.Params
	// 图片全称包括扩展名。目前支持 jpg jpeg png
	_pictureName string
	// 图片文件二进制流
	_img *model.File
	// true返回Https地址
	_isHttps bool
}

// NewTmallServicecenterPictureUploadRequest 初始化TmallServicecenterPictureUploadAPIRequest对象
func NewTmallServicecenterPictureUploadRequest() *TmallServicecenterPictureUploadAPIRequest {
	return &TmallServicecenterPictureUploadAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterPictureUploadAPIRequest) Reset() {
	r._pictureName = ""
	r._img = nil
	r._isHttps = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterPictureUploadAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterPictureUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterPictureUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureName is PictureName Setter
// 图片全称包括扩展名。目前支持 jpg jpeg png
func (r *TmallServicecenterPictureUploadAPIRequest) SetPictureName(_pictureName string) error {
	r._pictureName = _pictureName
	r.Set("picture_name", _pictureName)
	return nil
}

// GetPictureName PictureName Getter
func (r TmallServicecenterPictureUploadAPIRequest) GetPictureName() string {
	return r._pictureName
}

// SetImg is Img Setter
// 图片文件二进制流
func (r *TmallServicecenterPictureUploadAPIRequest) SetImg(_img *model.File) error {
	r._img = _img
	r.Set("img", _img)
	return nil
}

// GetImg Img Getter
func (r TmallServicecenterPictureUploadAPIRequest) GetImg() *model.File {
	return r._img
}

// SetIsHttps is IsHttps Setter
// true返回Https地址
func (r *TmallServicecenterPictureUploadAPIRequest) SetIsHttps(_isHttps bool) error {
	r._isHttps = _isHttps
	r.Set("is_https", _isHttps)
	return nil
}

// GetIsHttps IsHttps Getter
func (r TmallServicecenterPictureUploadAPIRequest) GetIsHttps() bool {
	return r._isHttps
}

var poolTmallServicecenterPictureUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterPictureUploadRequest()
	},
}

// GetTmallServicecenterPictureUploadRequest 从 sync.Pool 获取 TmallServicecenterPictureUploadAPIRequest
func GetTmallServicecenterPictureUploadAPIRequest() *TmallServicecenterPictureUploadAPIRequest {
	return poolTmallServicecenterPictureUploadAPIRequest.Get().(*TmallServicecenterPictureUploadAPIRequest)
}

// ReleaseTmallServicecenterPictureUploadAPIRequest 将 TmallServicecenterPictureUploadAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterPictureUploadAPIRequest(v *TmallServicecenterPictureUploadAPIRequest) {
	v.Reset()
	poolTmallServicecenterPictureUploadAPIRequest.Put(v)
}
