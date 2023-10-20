package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayimguploadAPIRequest 上传图片到支付宝图片空间接口 API请求
// taobao.alitrip.axin.trans.pay.img.upload
//
// 阿信供销平台-上传图片到支付宝图片空间接口
type TaobaoalitripaxintranspayimguploadAPIRequest struct {
	model.Params
	// 上传图片到支付宝图片空间接口入参
	_axinPayImgUploadDTO *AxinPayImgUploadDto
	// 图片字节流
	_imgContents *model.File
}

// NewTaobaoalitripaxintranspayimguploadRequest 初始化TaobaoalitripaxintranspayimguploadAPIRequest对象
func NewTaobaoalitripaxintranspayimguploadRequest() *TaobaoalitripaxintranspayimguploadAPIRequest {
	return &TaobaoalitripaxintranspayimguploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintranspayimguploadAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintranspayimguploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintranspayimguploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinPayImgUploadDTO is AxinPayImgUploadDTO Setter
// 上传图片到支付宝图片空间接口入参
func (r *TaobaoalitripaxintranspayimguploadAPIRequest) SetAxinPayImgUploadDTO(_axinPayImgUploadDTO *AxinPayImgUploadDto) error {
	r._axinPayImgUploadDTO = _axinPayImgUploadDTO
	r.Set("axin_pay_img_upload_d_t_o", _axinPayImgUploadDTO)
	return nil
}

// GetAxinPayImgUploadDTO AxinPayImgUploadDTO Getter
func (r TaobaoalitripaxintranspayimguploadAPIRequest) GetAxinPayImgUploadDTO() *AxinPayImgUploadDto {
	return r._axinPayImgUploadDTO
}

// SetImgContents is ImgContents Setter
// 图片字节流
func (r *TaobaoalitripaxintranspayimguploadAPIRequest) SetImgContents(_imgContents *model.File) error {
	r._imgContents = _imgContents
	r.Set("img_contents", _imgContents)
	return nil
}

// GetImgContents ImgContents Getter
func (r TaobaoalitripaxintranspayimguploadAPIRequest) GetImgContents() *model.File {
	return r._imgContents
}
