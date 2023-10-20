package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanmaterielgetAPIRequest 1-IBP同步物料接口 API请求
// alibaba.tmallgenie.scp.plan.materiel.get
//
// IBP同步物料接口
type AlibabatmallgeniescpplanmaterielgetAPIRequest struct {
	model.Params
	// 扩展字段
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplanmaterielgetRequest 初始化AlibabatmallgeniescpplanmaterielgetAPIRequest对象
func NewAlibabatmallgeniescpplanmaterielgetRequest() *AlibabatmallgeniescpplanmaterielgetAPIRequest {
	return &AlibabatmallgeniescpplanmaterielgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanmaterielgetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.materiel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanmaterielgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanmaterielgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展字段
func (r *AlibabatmallgeniescpplanmaterielgetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplanmaterielgetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
