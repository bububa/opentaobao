package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplansaleqtygetAPIRequest 12-同步销售数据 API请求
// alibaba.tmallgenie.scp.plan.sale.qty.get
//
// 同步销售数据
type AlibabatmallgeniescpplansaleqtygetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplansaleqtygetRequest 初始化AlibabatmallgeniescpplansaleqtygetAPIRequest对象
func NewAlibabatmallgeniescpplansaleqtygetRequest() *AlibabatmallgeniescpplansaleqtygetAPIRequest {
	return &AlibabatmallgeniescpplansaleqtygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplansaleqtygetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.sale.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplansaleqtygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplansaleqtygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabatmallgeniescpplansaleqtygetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplansaleqtygetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
