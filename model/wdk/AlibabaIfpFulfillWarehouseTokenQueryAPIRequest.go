package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIfpFulfillWarehouseTokenQueryAPIRequest 同城令牌打印接口 API请求
// alibaba.ifp.fulfill.warehouse.token.query
//
// 用于仓内作业打印包裹信息
type AlibabaIfpFulfillWarehouseTokenQueryAPIRequest struct {
	model.Params
	// 入参
	_packageQueryDTO *PackageQueryDto
}

// NewAlibabaIfpFulfillWarehouseTokenQueryRequest 初始化AlibabaIfpFulfillWarehouseTokenQueryAPIRequest对象
func NewAlibabaIfpFulfillWarehouseTokenQueryRequest() *AlibabaIfpFulfillWarehouseTokenQueryAPIRequest {
	return &AlibabaIfpFulfillWarehouseTokenQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) Reset() {
	r._packageQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ifp.fulfill.warehouse.token.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageQueryDTO is PackageQueryDTO Setter
// 入参
func (r *AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) SetPackageQueryDTO(_packageQueryDTO *PackageQueryDto) error {
	r._packageQueryDTO = _packageQueryDTO
	r.Set("package_query_d_t_o", _packageQueryDTO)
	return nil
}

// GetPackageQueryDTO PackageQueryDTO Getter
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetPackageQueryDTO() *PackageQueryDto {
	return r._packageQueryDTO
}

var poolAlibabaIfpFulfillWarehouseTokenQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIfpFulfillWarehouseTokenQueryRequest()
	},
}

// GetAlibabaIfpFulfillWarehouseTokenQueryRequest 从 sync.Pool 获取 AlibabaIfpFulfillWarehouseTokenQueryAPIRequest
func GetAlibabaIfpFulfillWarehouseTokenQueryAPIRequest() *AlibabaIfpFulfillWarehouseTokenQueryAPIRequest {
	return poolAlibabaIfpFulfillWarehouseTokenQueryAPIRequest.Get().(*AlibabaIfpFulfillWarehouseTokenQueryAPIRequest)
}

// ReleaseAlibabaIfpFulfillWarehouseTokenQueryAPIRequest 将 AlibabaIfpFulfillWarehouseTokenQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIfpFulfillWarehouseTokenQueryAPIRequest(v *AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) {
	v.Reset()
	poolAlibabaIfpFulfillWarehouseTokenQueryAPIRequest.Put(v)
}
