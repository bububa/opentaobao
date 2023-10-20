package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrgetupteminfoAPIRequest 获取上游温度信息（疫苗） API请求
// alibaba.alihealth.drug.kyt.dr.getupteminfo
//
// 根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
type AlibabaalihealthdrugkytdrgetupteminfoAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID
	_refEntId string
}

// NewAlibabaalihealthdrugkytdrgetupteminfoRequest 初始化AlibabaalihealthdrugkytdrgetupteminfoAPIRequest对象
func NewAlibabaalihealthdrugkytdrgetupteminfoRequest() *AlibabaalihealthdrugkytdrgetupteminfoAPIRequest {
	return &AlibabaalihealthdrugkytdrgetupteminfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getupteminfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrgetupteminfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}
