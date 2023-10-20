package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcreativehorizontalfindpageAPIRequest 横向管理创意分页查询 API请求
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
type TaobaouniversalbpcreativehorizontalfindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// creativeQueryVO
	_creativeQueryVO *CreativeQueryVo
}

// NewTaobaouniversalbpcreativehorizontalfindpageRequest 初始化TaobaouniversalbpcreativehorizontalfindpageAPIRequest对象
func NewTaobaouniversalbpcreativehorizontalfindpageRequest() *TaobaouniversalbpcreativehorizontalfindpageAPIRequest {
	return &TaobaouniversalbpcreativehorizontalfindpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcreativehorizontalfindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.creative.horizontal.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcreativehorizontalfindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcreativehorizontalfindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcreativehorizontalfindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcreativehorizontalfindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCreativeQueryVO is CreativeQueryVO Setter
// creativeQueryVO
func (r *TaobaouniversalbpcreativehorizontalfindpageAPIRequest) SetCreativeQueryVO(_creativeQueryVO *CreativeQueryVo) error {
	r._creativeQueryVO = _creativeQueryVO
	r.Set("creative_query_v_o", _creativeQueryVO)
	return nil
}

// GetCreativeQueryVO CreativeQueryVO Getter
func (r TaobaouniversalbpcreativehorizontalfindpageAPIRequest) GetCreativeQueryVO() *CreativeQueryVo {
	return r._creativeQueryVO
}
