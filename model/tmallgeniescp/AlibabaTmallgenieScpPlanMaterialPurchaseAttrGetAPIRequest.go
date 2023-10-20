package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest 物料的采购属性查询 API请求
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
type AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest struct {
	model.Params
	// 扩展字段
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest 初始化AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest() *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest {
	return &AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.material.purchase.attr.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展字段
func (r *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest()
	},
}

// GetAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest
func GetAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest() *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest {
	return poolAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest.Get().(*AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest 将 AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest(v *AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest.Put(v)
}
