package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMaterialAccessallowedAPIRequest 物料准入判断 API请求
// taobao.universalbp.material.accessallowed
//
// 入参推广信息，出参相关主体是否可投放。如果投放了风控不准入的商品，无法正常投放。
type TaobaoUniversalbpMaterialAccessallowedAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// materialAccessAllowQueryVO
	_materialAccessAllowQueryVO *MaterialAccessAllowQueryVo
}

// NewTaobaoUniversalbpMaterialAccessallowedRequest 初始化TaobaoUniversalbpMaterialAccessallowedAPIRequest对象
func NewTaobaoUniversalbpMaterialAccessallowedRequest() *TaobaoUniversalbpMaterialAccessallowedAPIRequest {
	return &TaobaoUniversalbpMaterialAccessallowedAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpMaterialAccessallowedAPIRequest) Reset() {
	r._topServiceContext = nil
	r._materialAccessAllowQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpMaterialAccessallowedAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.material.accessallowed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpMaterialAccessallowedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpMaterialAccessallowedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpMaterialAccessallowedAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpMaterialAccessallowedAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetMaterialAccessAllowQueryVO is MaterialAccessAllowQueryVO Setter
// materialAccessAllowQueryVO
func (r *TaobaoUniversalbpMaterialAccessallowedAPIRequest) SetMaterialAccessAllowQueryVO(_materialAccessAllowQueryVO *MaterialAccessAllowQueryVo) error {
	r._materialAccessAllowQueryVO = _materialAccessAllowQueryVO
	r.Set("material_access_allow_query_v_o", _materialAccessAllowQueryVO)
	return nil
}

// GetMaterialAccessAllowQueryVO MaterialAccessAllowQueryVO Getter
func (r TaobaoUniversalbpMaterialAccessallowedAPIRequest) GetMaterialAccessAllowQueryVO() *MaterialAccessAllowQueryVo {
	return r._materialAccessAllowQueryVO
}

var poolTaobaoUniversalbpMaterialAccessallowedAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpMaterialAccessallowedRequest()
	},
}

// GetTaobaoUniversalbpMaterialAccessallowedRequest 从 sync.Pool 获取 TaobaoUniversalbpMaterialAccessallowedAPIRequest
func GetTaobaoUniversalbpMaterialAccessallowedAPIRequest() *TaobaoUniversalbpMaterialAccessallowedAPIRequest {
	return poolTaobaoUniversalbpMaterialAccessallowedAPIRequest.Get().(*TaobaoUniversalbpMaterialAccessallowedAPIRequest)
}

// ReleaseTaobaoUniversalbpMaterialAccessallowedAPIRequest 将 TaobaoUniversalbpMaterialAccessallowedAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpMaterialAccessallowedAPIRequest(v *TaobaoUniversalbpMaterialAccessallowedAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpMaterialAccessallowedAPIRequest.Put(v)
}
