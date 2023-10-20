package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest 码的药品信息查询 API请求
// alibaba.alihealth.drug.kyt.getcodebaseinfo
//
// 提供根据码查询码基本信息接口
type AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 码
	_code string
}

// NewAlibabaalihealthdrugkytgetcodebaseinfoRequest 初始化AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest对象
func NewAlibabaalihealthdrugkytgetcodebaseinfoRequest() *AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest {
	return &AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getcodebaseinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 码
func (r *AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest) GetCode() string {
	return r._code
}
