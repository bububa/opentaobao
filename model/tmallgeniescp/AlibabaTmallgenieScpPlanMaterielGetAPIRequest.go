package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanMaterielGetAPIRequest 1-IBP同步物料接口 API请求
// alibaba.tmallgenie.scp.plan.materiel.get
//
// IBP同步物料接口
type AlibabaTmallgenieScpPlanMaterielGetAPIRequest struct {
	model.Params
	// 扩展字段
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanMaterielGetRequest 初始化AlibabaTmallgenieScpPlanMaterielGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanMaterielGetRequest() *AlibabaTmallgenieScpPlanMaterielGetAPIRequest {
	return &AlibabaTmallgenieScpPlanMaterielGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.materiel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestExtendJson Setter
// 扩展字段
func (r *AlibabaTmallgenieScpPlanMaterielGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// Get RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
