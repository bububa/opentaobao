package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemhscodeauditresultsqueryAPIRequest 商品hscode信息审核状态查询接口 API请求
// tmall.item.hscode.audit.results.query
//
// 通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
type TmallitemhscodeauditresultsqueryAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallitemhscodeauditresultsqueryRequest 初始化TmallitemhscodeauditresultsqueryAPIRequest对象
func NewTmallitemhscodeauditresultsqueryRequest() *TmallitemhscodeauditresultsqueryAPIRequest {
	return &TmallitemhscodeauditresultsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemhscodeauditresultsqueryAPIRequest) GetApiMethodName() string {
	return "tmall.item.hscode.audit.results.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemhscodeauditresultsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemhscodeauditresultsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemhscodeauditresultsqueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemhscodeauditresultsqueryAPIRequest) GetItemId() int64 {
	return r._itemId
}
