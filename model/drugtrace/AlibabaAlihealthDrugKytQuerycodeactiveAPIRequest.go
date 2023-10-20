package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytquerycodeactiveAPIRequest 查询码是否激活 API请求
// alibaba.alihealth.drug.kyt.querycodeactive
//
// 查询码是否激活
type AlibabaalihealthdrugkytquerycodeactiveAPIRequest struct {
	model.Params
	// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaalihealthdrugkytquerycodeactiveRequest 初始化AlibabaalihealthdrugkytquerycodeactiveAPIRequest对象
func NewAlibabaalihealthdrugkytquerycodeactiveRequest() *AlibabaalihealthdrugkytquerycodeactiveAPIRequest {
	return &AlibabaalihealthdrugkytquerycodeactiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytquerycodeactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.querycodeactive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytquerycodeactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytquerycodeactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaalihealthdrugkytquerycodeactiveAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugkytquerycodeactiveAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaalihealthdrugkytquerycodeactiveAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytquerycodeactiveAPIRequest) GetRefEntId() string {
	return r._refEntId
}
