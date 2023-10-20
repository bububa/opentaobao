package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpmaterialaccessallowedAPIRequest 物料准入判断 API请求
// taobao.universalbp.material.accessallowed
//
// 入参推广信息，出参相关主体是否可投放。如果投放了风控不准入的商品，无法正常投放。
type TaobaouniversalbpmaterialaccessallowedAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// materialAccessAllowQueryVO
	_materialAccessAllowQueryVO *MaterialAccessAllowQueryVo
}

// NewTaobaouniversalbpmaterialaccessallowedRequest 初始化TaobaouniversalbpmaterialaccessallowedAPIRequest对象
func NewTaobaouniversalbpmaterialaccessallowedRequest() *TaobaouniversalbpmaterialaccessallowedAPIRequest {
	return &TaobaouniversalbpmaterialaccessallowedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpmaterialaccessallowedAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.material.accessallowed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpmaterialaccessallowedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpmaterialaccessallowedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpmaterialaccessallowedAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpmaterialaccessallowedAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetMaterialAccessAllowQueryVO is MaterialAccessAllowQueryVO Setter
// materialAccessAllowQueryVO
func (r *TaobaouniversalbpmaterialaccessallowedAPIRequest) SetMaterialAccessAllowQueryVO(_materialAccessAllowQueryVO *MaterialAccessAllowQueryVo) error {
	r._materialAccessAllowQueryVO = _materialAccessAllowQueryVO
	r.Set("material_access_allow_query_v_o", _materialAccessAllowQueryVO)
	return nil
}

// GetMaterialAccessAllowQueryVO MaterialAccessAllowQueryVO Getter
func (r TaobaouniversalbpmaterialaccessallowedAPIRequest) GetMaterialAccessAllowQueryVO() *MaterialAccessAllowQueryVo {
	return r._materialAccessAllowQueryVO
}
