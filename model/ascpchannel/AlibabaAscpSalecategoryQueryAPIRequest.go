package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpsalecategoryqueryAPIRequest 货品品类查询 API请求
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
type AlibabaascpsalecategoryqueryAPIRequest struct {
	model.Params
	// 货品ID
	_itemId []int64
}

// NewAlibabaascpsalecategoryqueryRequest 初始化AlibabaascpsalecategoryqueryAPIRequest对象
func NewAlibabaascpsalecategoryqueryRequest() *AlibabaascpsalecategoryqueryAPIRequest {
	return &AlibabaascpsalecategoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpsalecategoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.salecategory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpsalecategoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpsalecategoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 货品ID
func (r *AlibabaascpsalecategoryqueryAPIRequest) SetItemId(_itemId []int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaascpsalecategoryqueryAPIRequest) GetItemId() []int64 {
	return r._itemId
}
