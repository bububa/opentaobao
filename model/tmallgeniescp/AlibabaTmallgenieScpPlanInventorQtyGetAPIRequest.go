package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplaninventorqtygetAPIRequest 10-同步库存现有量 API请求
// alibaba.tmallgenie.scp.plan.inventor.qty.get
//
// 同步库存现有量
type AlibabatmallgeniescpplaninventorqtygetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplaninventorqtygetRequest 初始化AlibabatmallgeniescpplaninventorqtygetAPIRequest对象
func NewAlibabatmallgeniescpplaninventorqtygetRequest() *AlibabatmallgeniescpplaninventorqtygetAPIRequest {
	return &AlibabatmallgeniescpplaninventorqtygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplaninventorqtygetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.inventor.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplaninventorqtygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplaninventorqtygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabatmallgeniescpplaninventorqtygetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplaninventorqtygetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
