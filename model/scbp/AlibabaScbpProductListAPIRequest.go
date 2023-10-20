package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductListAPIRequest 查询P4P产品 API请求
// alibaba.scbp.product.list
//
// 查询P4P产品
type AlibabaScbpProductListAPIRequest struct {
	model.Params
	// 产品分组标识
	_groupId string
	// 产品分页查询，每页个数，最大值20
	_perPageSize int64
	// 第几页
	_toPage int64
}

// NewAlibabaScbpProductListRequest 初始化AlibabaScbpProductListAPIRequest对象
func NewAlibabaScbpProductListRequest() *AlibabaScbpProductListAPIRequest {
	return &AlibabaScbpProductListAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpProductListAPIRequest) Reset() {
	r._groupId = ""
	r._perPageSize = 0
	r._toPage = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpProductListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 产品分组标识
func (r *AlibabaScbpProductListAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaScbpProductListAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetPerPageSize is PerPageSize Setter
// 产品分页查询，每页个数，最大值20
func (r *AlibabaScbpProductListAPIRequest) SetPerPageSize(_perPageSize int64) error {
	r._perPageSize = _perPageSize
	r.Set("per_page_size", _perPageSize)
	return nil
}

// GetPerPageSize PerPageSize Getter
func (r AlibabaScbpProductListAPIRequest) GetPerPageSize() int64 {
	return r._perPageSize
}

// SetToPage is ToPage Setter
// 第几页
func (r *AlibabaScbpProductListAPIRequest) SetToPage(_toPage int64) error {
	r._toPage = _toPage
	r.Set("to_page", _toPage)
	return nil
}

// GetToPage ToPage Getter
func (r AlibabaScbpProductListAPIRequest) GetToPage() int64 {
	return r._toPage
}

var poolAlibabaScbpProductListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpProductListRequest()
	},
}

// GetAlibabaScbpProductListRequest 从 sync.Pool 获取 AlibabaScbpProductListAPIRequest
func GetAlibabaScbpProductListAPIRequest() *AlibabaScbpProductListAPIRequest {
	return poolAlibabaScbpProductListAPIRequest.Get().(*AlibabaScbpProductListAPIRequest)
}

// ReleaseAlibabaScbpProductListAPIRequest 将 AlibabaScbpProductListAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpProductListAPIRequest(v *AlibabaScbpProductListAPIRequest) {
	v.Reset()
	poolAlibabaScbpProductListAPIRequest.Put(v)
}
