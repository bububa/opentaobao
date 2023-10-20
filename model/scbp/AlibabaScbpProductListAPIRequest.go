package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpproductlistAPIRequest 查询P4P产品 API请求
// alibaba.scbp.product.list
//
// 查询P4P产品
type AlibabascbpproductlistAPIRequest struct {
	model.Params
	// 产品分组标识
	_groupId string
	// 产品分页查询，每页个数，最大值20
	_perPageSize int64
	// 第几页
	_toPage int64
}

// NewAlibabascbpproductlistRequest 初始化AlibabascbpproductlistAPIRequest对象
func NewAlibabascbpproductlistRequest() *AlibabascbpproductlistAPIRequest {
	return &AlibabascbpproductlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpproductlistAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpproductlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpproductlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 产品分组标识
func (r *AlibabascbpproductlistAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabascbpproductlistAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetPerPageSize is PerPageSize Setter
// 产品分页查询，每页个数，最大值20
func (r *AlibabascbpproductlistAPIRequest) SetPerPageSize(_perPageSize int64) error {
	r._perPageSize = _perPageSize
	r.Set("per_page_size", _perPageSize)
	return nil
}

// GetPerPageSize PerPageSize Getter
func (r AlibabascbpproductlistAPIRequest) GetPerPageSize() int64 {
	return r._perPageSize
}

// SetToPage is ToPage Setter
// 第几页
func (r *AlibabascbpproductlistAPIRequest) SetToPage(_toPage int64) error {
	r._toPage = _toPage
	r.Set("to_page", _toPage)
	return nil
}

// GetToPage ToPage Getter
func (r AlibabascbpproductlistAPIRequest) GetToPage() int64 {
	return r._toPage
}
