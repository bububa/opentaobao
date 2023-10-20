package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpSalecategoryQueryAPIRequest 货品品类查询 API请求
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
type AlibabaAscpSalecategoryQueryAPIRequest struct {
	model.Params
	// 货品ID
	_itemId []int64
}

// NewAlibabaAscpSalecategoryQueryRequest 初始化AlibabaAscpSalecategoryQueryAPIRequest对象
func NewAlibabaAscpSalecategoryQueryRequest() *AlibabaAscpSalecategoryQueryAPIRequest {
	return &AlibabaAscpSalecategoryQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpSalecategoryQueryAPIRequest) Reset() {
	r._itemId = r._itemId[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.salecategory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 货品ID
func (r *AlibabaAscpSalecategoryQueryAPIRequest) SetItemId(_itemId []int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetItemId() []int64 {
	return r._itemId
}

var poolAlibabaAscpSalecategoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpSalecategoryQueryRequest()
	},
}

// GetAlibabaAscpSalecategoryQueryRequest 从 sync.Pool 获取 AlibabaAscpSalecategoryQueryAPIRequest
func GetAlibabaAscpSalecategoryQueryAPIRequest() *AlibabaAscpSalecategoryQueryAPIRequest {
	return poolAlibabaAscpSalecategoryQueryAPIRequest.Get().(*AlibabaAscpSalecategoryQueryAPIRequest)
}

// ReleaseAlibabaAscpSalecategoryQueryAPIRequest 将 AlibabaAscpSalecategoryQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpSalecategoryQueryAPIRequest(v *AlibabaAscpSalecategoryQueryAPIRequest) {
	v.Reset()
	poolAlibabaAscpSalecategoryQueryAPIRequest.Put(v)
}
