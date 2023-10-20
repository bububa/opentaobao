package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytsaveentAPIRequest 新增往来单位企业 API请求
// alibaba.alihealth.drug.kyt.saveent
//
// 新增往来单位企业记录
type AlibabaalihealthdrugkytsaveentAPIRequest struct {
	model.Params
	// 添加企业唯一标识
	_refEntId string
	// 新增企业信息
	_addEntReq *AddEntReqDto
	// 图片数据流。图片大小务必控制在2M以内
	_licPictureByte *model.File
}

// NewAlibabaalihealthdrugkytsaveentRequest 初始化AlibabaalihealthdrugkytsaveentAPIRequest对象
func NewAlibabaalihealthdrugkytsaveentRequest() *AlibabaalihealthdrugkytsaveentAPIRequest {
	return &AlibabaalihealthdrugkytsaveentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytsaveentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.saveent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytsaveentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytsaveentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 添加企业唯一标识
func (r *AlibabaalihealthdrugkytsaveentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytsaveentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetAddEntReq is AddEntReq Setter
// 新增企业信息
func (r *AlibabaalihealthdrugkytsaveentAPIRequest) SetAddEntReq(_addEntReq *AddEntReqDto) error {
	r._addEntReq = _addEntReq
	r.Set("add_ent_req", _addEntReq)
	return nil
}

// GetAddEntReq AddEntReq Getter
func (r AlibabaalihealthdrugkytsaveentAPIRequest) GetAddEntReq() *AddEntReqDto {
	return r._addEntReq
}

// SetLicPictureByte is LicPictureByte Setter
// 图片数据流。图片大小务必控制在2M以内
func (r *AlibabaalihealthdrugkytsaveentAPIRequest) SetLicPictureByte(_licPictureByte *model.File) error {
	r._licPictureByte = _licPictureByte
	r.Set("lic_picture_byte", _licPictureByte)
	return nil
}

// GetLicPictureByte LicPictureByte Getter
func (r AlibabaalihealthdrugkytsaveentAPIRequest) GetLicPictureByte() *model.File {
	return r._licPictureByte
}
