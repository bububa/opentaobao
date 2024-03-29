package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductStatusUpdateAPIRequest 修改P4P产品推广状态 API请求
// alibaba.scbp.product.status.update
//
// 修改P4P产品推广状态
type AlibabaScbpProductStatusUpdateAPIRequest struct {
	model.Params
	// 产品ID列表
	_productIdList []string
	// enabled:开启,disabled:暂停
	_status string
}

// NewAlibabaScbpProductStatusUpdateRequest 初始化AlibabaScbpProductStatusUpdateAPIRequest对象
func NewAlibabaScbpProductStatusUpdateRequest() *AlibabaScbpProductStatusUpdateAPIRequest {
	return &AlibabaScbpProductStatusUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpProductStatusUpdateAPIRequest) Reset() {
	r._productIdList = r._productIdList[:0]
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIdList is ProductIdList Setter
// 产品ID列表
func (r *AlibabaScbpProductStatusUpdateAPIRequest) SetProductIdList(_productIdList []string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// GetProductIdList ProductIdList Getter
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetProductIdList() []string {
	return r._productIdList
}

// SetStatus is Status Setter
// enabled:开启,disabled:暂停
func (r *AlibabaScbpProductStatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetStatus() string {
	return r._status
}

var poolAlibabaScbpProductStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpProductStatusUpdateRequest()
	},
}

// GetAlibabaScbpProductStatusUpdateRequest 从 sync.Pool 获取 AlibabaScbpProductStatusUpdateAPIRequest
func GetAlibabaScbpProductStatusUpdateAPIRequest() *AlibabaScbpProductStatusUpdateAPIRequest {
	return poolAlibabaScbpProductStatusUpdateAPIRequest.Get().(*AlibabaScbpProductStatusUpdateAPIRequest)
}

// ReleaseAlibabaScbpProductStatusUpdateAPIRequest 将 AlibabaScbpProductStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpProductStatusUpdateAPIRequest(v *AlibabaScbpProductStatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaScbpProductStatusUpdateAPIRequest.Put(v)
}
