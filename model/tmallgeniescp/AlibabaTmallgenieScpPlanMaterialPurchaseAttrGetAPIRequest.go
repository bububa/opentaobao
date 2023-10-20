package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest 物料的采购属性查询 API请求
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
type AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest struct {
	model.Params
	// 扩展字段
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplanmaterialpurchaseattrgetRequest 初始化AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest对象
func NewAlibabatmallgeniescpplanmaterialpurchaseattrgetRequest() *AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest {
	return &AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.material.purchase.attr.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展字段
func (r *AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplanmaterialpurchaseattrgetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
