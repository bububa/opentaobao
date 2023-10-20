package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinpovupdateAPIRequest 新增/变更接种点信息 API请求
// alibaba.health.vaccin.pov.update
//
// ISV 将疫苗的接种点信息同步到免疫规划中心，提醒用户接种时可提供接种点详情。
type AlibabahealthvaccinpovupdateAPIRequest struct {
	model.Params
	// 接种点编码
	_povNo string
	// 接种点名称
	_povName string
	// 接种点联系电话
	_telephone string
	// 接种点具体地址
	_address string
	// 接种点介绍
	_description string
	// 服务时间
	_businessTime string
}

// NewAlibabahealthvaccinpovupdateRequest 初始化AlibabahealthvaccinpovupdateAPIRequest对象
func NewAlibabahealthvaccinpovupdateRequest() *AlibabahealthvaccinpovupdateAPIRequest {
	return &AlibabahealthvaccinpovupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinpovupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.pov.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinpovupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinpovupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPovNo is PovNo Setter
// 接种点编码
func (r *AlibabahealthvaccinpovupdateAPIRequest) SetPovNo(_povNo string) error {
	r._povNo = _povNo
	r.Set("pov_no", _povNo)
	return nil
}

// GetPovNo PovNo Getter
func (r AlibabahealthvaccinpovupdateAPIRequest) GetPovNo() string {
	return r._povNo
}

// SetPovName is PovName Setter
// 接种点名称
func (r *AlibabahealthvaccinpovupdateAPIRequest) SetPovName(_povName string) error {
	r._povName = _povName
	r.Set("pov_name", _povName)
	return nil
}

// GetPovName PovName Getter
func (r AlibabahealthvaccinpovupdateAPIRequest) GetPovName() string {
	return r._povName
}

// SetTelephone is Telephone Setter
// 接种点联系电话
func (r *AlibabahealthvaccinpovupdateAPIRequest) SetTelephone(_telephone string) error {
	r._telephone = _telephone
	r.Set("telephone", _telephone)
	return nil
}

// GetTelephone Telephone Getter
func (r AlibabahealthvaccinpovupdateAPIRequest) GetTelephone() string {
	return r._telephone
}

// SetAddress is Address Setter
// 接种点具体地址
func (r *AlibabahealthvaccinpovupdateAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabahealthvaccinpovupdateAPIRequest) GetAddress() string {
	return r._address
}

// SetDescription is Description Setter
// 接种点介绍
func (r *AlibabahealthvaccinpovupdateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r AlibabahealthvaccinpovupdateAPIRequest) GetDescription() string {
	return r._description
}

// SetBusinessTime is BusinessTime Setter
// 服务时间
func (r *AlibabahealthvaccinpovupdateAPIRequest) SetBusinessTime(_businessTime string) error {
	r._businessTime = _businessTime
	r.Set("business_time", _businessTime)
	return nil
}

// GetBusinessTime BusinessTime Getter
func (r AlibabahealthvaccinpovupdateAPIRequest) GetBusinessTime() string {
	return r._businessTime
}
