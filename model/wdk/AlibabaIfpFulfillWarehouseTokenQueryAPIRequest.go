package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ifp.fulfill.warehouse.token.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PackageQueryDTO Setter
// 入参
func (r *AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) SetPackageQueryDTO(_packageQueryDTO *PackageQueryDto) error {
	r._packageQueryDTO = _packageQueryDTO
	r.Set("package_query_d_t_o", _packageQueryDTO)
	return nil
}

// Get PackageQueryDTO Getter
func (r AlibabaIfpFulfillWarehouseTokenQueryAPIRequest) GetPackageQueryDTO() *PackageQueryDto {
	return r._packageQueryDTO
}
