package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpmaterialitemfindpageAPIRequest 分页查询商品信息 API请求
// taobao.universalbp.material.item.findpage
//
// 分页获取店铺内的商品列表
type TaobaouniversalbpmaterialitemfindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// itemQueryVO
	_itemQueryVO *ItemQueryVo
}

// NewTaobaouniversalbpmaterialitemfindpageRequest 初始化TaobaouniversalbpmaterialitemfindpageAPIRequest对象
func NewTaobaouniversalbpmaterialitemfindpageRequest() *TaobaouniversalbpmaterialitemfindpageAPIRequest {
	return &TaobaouniversalbpmaterialitemfindpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpmaterialitemfindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.material.item.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpmaterialitemfindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpmaterialitemfindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpmaterialitemfindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpmaterialitemfindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetItemQueryVO is ItemQueryVO Setter
// itemQueryVO
func (r *TaobaouniversalbpmaterialitemfindpageAPIRequest) SetItemQueryVO(_itemQueryVO *ItemQueryVo) error {
	r._itemQueryVO = _itemQueryVO
	r.Set("item_query_v_o", _itemQueryVO)
	return nil
}

// GetItemQueryVO ItemQueryVO Getter
func (r TaobaouniversalbpmaterialitemfindpageAPIRequest) GetItemQueryVO() *ItemQueryVo {
	return r._itemQueryVO
}
