package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest 扫码营销码查询 API请求
// alibaba.alihealth.drug.code.kyt.smyx.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaalihealthdrugcodekytsmyxquerycodeRequest 初始化AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest对象
func NewAlibabaalihealthdrugcodekytsmyxquerycodeRequest() *AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest {
	return &AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.smyx.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
