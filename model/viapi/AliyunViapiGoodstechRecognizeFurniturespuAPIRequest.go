package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapigoodstechrecognizefurniturespuAPIRequest 家居SPU识别 API请求
// aliyun.viapi.goodstech.recognize.furniturespu
//
// 对输入的家居模型图进行分类，目前类别数可达70类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapigoodstechrecognizefurniturespuAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
	// 模型x方向的尺寸，单位cm，默认值100
	_xLength string
	// 模型y方向的尺寸，单位cm，默认值100
	_yLength string
	// 模型z方向的尺寸，单位cm，默认值100
	_zLength string
}

// NewAliyunviapigoodstechrecognizefurniturespuRequest 初始化AliyunviapigoodstechrecognizefurniturespuAPIRequest对象
func NewAliyunviapigoodstechrecognizefurniturespuRequest() *AliyunviapigoodstechrecognizefurniturespuAPIRequest {
	return &AliyunviapigoodstechrecognizefurniturespuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.goodstech.recognize.furniturespu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunviapigoodstechrecognizefurniturespuAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

// SetXLength is XLength Setter
// 模型x方向的尺寸，单位cm，默认值100
func (r *AliyunviapigoodstechrecognizefurniturespuAPIRequest) SetXLength(_xLength string) error {
	r._xLength = _xLength
	r.Set("x_length", _xLength)
	return nil
}

// GetXLength XLength Getter
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetXLength() string {
	return r._xLength
}

// SetYLength is YLength Setter
// 模型y方向的尺寸，单位cm，默认值100
func (r *AliyunviapigoodstechrecognizefurniturespuAPIRequest) SetYLength(_yLength string) error {
	r._yLength = _yLength
	r.Set("y_length", _yLength)
	return nil
}

// GetYLength YLength Getter
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetYLength() string {
	return r._yLength
}

// SetZLength is ZLength Setter
// 模型z方向的尺寸，单位cm，默认值100
func (r *AliyunviapigoodstechrecognizefurniturespuAPIRequest) SetZLength(_zLength string) error {
	r._zLength = _zLength
	r.Set("z_length", _zLength)
	return nil
}

// GetZLength ZLength Getter
func (r AliyunviapigoodstechrecognizefurniturespuAPIRequest) GetZLength() string {
	return r._zLength
}
