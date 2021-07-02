package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemHscodeAuditResultsQueryAPIRequest 商品hscode信息审核状态查询接口 API请求
// tmall.item.hscode.audit.results.query
//
// 通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
type TmallItemHscodeAuditResultsQueryAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallItemHscodeAuditResultsQueryRequest 初始化TmallItemHscodeAuditResultsQueryAPIRequest对象
func NewTmallItemHscodeAuditResultsQueryRequest() *TmallItemHscodeAuditResultsQueryAPIRequest {
	return &TmallItemHscodeAuditResultsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemHscodeAuditResultsQueryAPIRequest) GetApiMethodName() string {
	return "tmall.item.hscode.audit.results.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemHscodeAuditResultsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemHscodeAuditResultsQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemHscodeAuditResultsQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}
