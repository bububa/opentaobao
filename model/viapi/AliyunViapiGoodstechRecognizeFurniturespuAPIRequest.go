package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiGoodstechRecognizeFurniturespuAPIRequest 家居SPU识别 API请求
// aliyun.viapi.goodstech.recognize.furniturespu
//
// 对输入的家居模型图进行分类，目前类别数可达70类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiGoodstechRecognizeFurniturespuAPIRequest struct {
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

// NewAliyunViapiGoodstechRecognizeFurniturespuRequest 初始化AliyunViapiGoodstechRecognizeFurniturespuAPIRequest对象
func NewAliyunViapiGoodstechRecognizeFurniturespuRequest() *AliyunViapiGoodstechRecognizeFurniturespuAPIRequest {
	return &AliyunViapiGoodstechRecognizeFurniturespuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.goodstech.recognize.furniturespu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

// SetXLength is XLength Setter
// 模型x方向的尺寸，单位cm，默认值100
func (r *AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) SetXLength(_xLength string) error {
	r._xLength = _xLength
	r.Set("x_length", _xLength)
	return nil
}

// GetXLength XLength Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) GetXLength() string {
	return r._xLength
}

// SetYLength is YLength Setter
// 模型y方向的尺寸，单位cm，默认值100
func (r *AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) SetYLength(_yLength string) error {
	r._yLength = _yLength
	r.Set("y_length", _yLength)
	return nil
}

// GetYLength YLength Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) GetYLength() string {
	return r._yLength
}

// SetZLength is ZLength Setter
// 模型z方向的尺寸，单位cm，默认值100
func (r *AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) SetZLength(_zLength string) error {
	r._zLength = _zLength
	r.Set("z_length", _zLength)
	return nil
}

// GetZLength ZLength Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuAPIRequest) GetZLength() string {
	return r._zLength
}
