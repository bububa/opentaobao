package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosBrandCoproductGroupUserCountAPIRequest 按照查询条件统计总数 API请求
// alibaba.mos.brand.coproduct.group.user.count
//
// 按照查询条件统计总数
type AlibabaMosBrandCoproductGroupUserCountAPIRequest struct {
	model.Params
	// 查询条件
	_brandCoProductGroupUserQueryParam *BrandCoProductGroupUserQueryParam
}

// NewAlibabaMosBrandCoproductGroupUserCountRequest 初始化AlibabaMosBrandCoproductGroupUserCountAPIRequest对象
func NewAlibabaMosBrandCoproductGroupUserCountRequest() *AlibabaMosBrandCoproductGroupUserCountAPIRequest {
	return &AlibabaMosBrandCoproductGroupUserCountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosBrandCoproductGroupUserCountAPIRequest) Reset() {
	r._brandCoProductGroupUserQueryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosBrandCoproductGroupUserCountAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.brand.coproduct.group.user.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosBrandCoproductGroupUserCountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosBrandCoproductGroupUserCountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandCoProductGroupUserQueryParam is BrandCoProductGroupUserQueryParam Setter
// 查询条件
func (r *AlibabaMosBrandCoproductGroupUserCountAPIRequest) SetBrandCoProductGroupUserQueryParam(_brandCoProductGroupUserQueryParam *BrandCoProductGroupUserQueryParam) error {
	r._brandCoProductGroupUserQueryParam = _brandCoProductGroupUserQueryParam
	r.Set("brand_co_product_group_user_query_param", _brandCoProductGroupUserQueryParam)
	return nil
}

// GetBrandCoProductGroupUserQueryParam BrandCoProductGroupUserQueryParam Getter
func (r AlibabaMosBrandCoproductGroupUserCountAPIRequest) GetBrandCoProductGroupUserQueryParam() *BrandCoProductGroupUserQueryParam {
	return r._brandCoProductGroupUserQueryParam
}

var poolAlibabaMosBrandCoproductGroupUserCountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosBrandCoproductGroupUserCountRequest()
	},
}

// GetAlibabaMosBrandCoproductGroupUserCountRequest 从 sync.Pool 获取 AlibabaMosBrandCoproductGroupUserCountAPIRequest
func GetAlibabaMosBrandCoproductGroupUserCountAPIRequest() *AlibabaMosBrandCoproductGroupUserCountAPIRequest {
	return poolAlibabaMosBrandCoproductGroupUserCountAPIRequest.Get().(*AlibabaMosBrandCoproductGroupUserCountAPIRequest)
}

// ReleaseAlibabaMosBrandCoproductGroupUserCountAPIRequest 将 AlibabaMosBrandCoproductGroupUserCountAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosBrandCoproductGroupUserCountAPIRequest(v *AlibabaMosBrandCoproductGroupUserCountAPIRequest) {
	v.Reset()
	poolAlibabaMosBrandCoproductGroupUserCountAPIRequest.Put(v)
}
