package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城令牌打印接口 API请求
alibaba.ifp.fulfill.warehouse.token.query

用于仓内作业打印包裹信息
*/
type AlibabaIfpFulfillWarehouseTokenQueryRequest struct {
    model.Params
    // 入参
    _packageQueryDTO   *PackageQueryDto
}

// 初始化AlibabaIfpFulfillWarehouseTokenQueryRequest对象
func NewAlibabaIfpFulfillWarehouseTokenQueryRequest() *AlibabaIfpFulfillWarehouseTokenQueryRequest{
    return &AlibabaIfpFulfillWarehouseTokenQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIfpFulfillWarehouseTokenQueryRequest) GetApiMethodName() string {
    return "alibaba.ifp.fulfill.warehouse.token.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIfpFulfillWarehouseTokenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PackageQueryDTO Setter
// 入参
func (r *AlibabaIfpFulfillWarehouseTokenQueryRequest) SetPackageQueryDTO(_packageQueryDTO *PackageQueryDto) error {
    r._packageQueryDTO = _packageQueryDTO
    r.Set("package_query_d_t_o", _packageQueryDTO)
    return nil
}

// PackageQueryDTO Getter
func (r AlibabaIfpFulfillWarehouseTokenQueryRequest) GetPackageQueryDTO() *PackageQueryDto {
    return r._packageQueryDTO
}
