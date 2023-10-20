package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSaveentAPIRequest 新增往来单位企业记录 API请求
// alibaba.alihealth.drug.kyt.wes.saveent
//
// 新增往来单位企业记录
type AlibabaAlihealthDrugKytWesSaveentAPIRequest struct {
	model.Params
	// 添加企业唯一标识
	_refEntId string
	// 服务时间
	_licenseToken string
	// 新增企业信息
	_addEntReq *AddEntReqDto
	// 图片数据流。图片大小务必控制在2M以内
	_licPictureByte *model.File
}

// NewAlibabaAlihealthDrugKytWesSaveentRequest 初始化AlibabaAlihealthDrugKytWesSaveentAPIRequest对象
func NewAlibabaAlihealthDrugKytWesSaveentRequest() *AlibabaAlihealthDrugKytWesSaveentAPIRequest {
	return &AlibabaAlihealthDrugKytWesSaveentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.saveent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 添加企业唯一标识
func (r *AlibabaAlihealthDrugKytWesSaveentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesSaveentAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetAddEntReq is AddEntReq Setter
// 新增企业信息
func (r *AlibabaAlihealthDrugKytWesSaveentAPIRequest) SetAddEntReq(_addEntReq *AddEntReqDto) error {
	r._addEntReq = _addEntReq
	r.Set("add_ent_req", _addEntReq)
	return nil
}

// GetAddEntReq AddEntReq Getter
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetAddEntReq() *AddEntReqDto {
	return r._addEntReq
}

// SetLicPictureByte is LicPictureByte Setter
// 图片数据流。图片大小务必控制在2M以内
func (r *AlibabaAlihealthDrugKytWesSaveentAPIRequest) SetLicPictureByte(_licPictureByte *model.File) error {
	r._licPictureByte = _licPictureByte
	r.Set("lic_picture_byte", _licPictureByte)
	return nil
}

// GetLicPictureByte LicPictureByte Getter
func (r AlibabaAlihealthDrugKytWesSaveentAPIRequest) GetLicPictureByte() *model.File {
	return r._licPictureByte
}
