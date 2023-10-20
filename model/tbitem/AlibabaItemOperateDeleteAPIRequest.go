package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateDeleteAPIRequest 商品删除 API请求
// alibaba.item.operate.delete
//
// 商品删除
type AlibabaItemOperateDeleteAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewAlibabaItemOperateDeleteRequest 初始化AlibabaItemOperateDeleteAPIRequest对象
func NewAlibabaItemOperateDeleteRequest() *AlibabaItemOperateDeleteAPIRequest {
	return &AlibabaItemOperateDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemOperateDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.item.operate.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemOperateDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemOperateDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemOperateDeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemOperateDeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}
