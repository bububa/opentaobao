package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城令牌打印接口 APIRequest
alibaba.ifp.fulfill.warehouse.token.query

用于仓内作业打印包裹信息
*/
type AlibabaIfpFulfillWarehouseTokenQueryRequest struct {
    model.Params

    // 入参
    packageQueryDTO   *PackageQueryDto 

}

func NewAlibabaIfpFulfillWarehouseTokenQueryRequest() *AlibabaIfpFulfillWarehouseTokenQueryRequest{
    return &AlibabaIfpFulfillWarehouseTokenQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIfpFulfillWarehouseTokenQueryRequest) GetApiMethodName() string {
    return "alibaba.ifp.fulfill.warehouse.token.query"
}

func (r AlibabaIfpFulfillWarehouseTokenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIfpFulfillWarehouseTokenQueryRequest) SetPackageQueryDTO(packageQueryDTO *PackageQueryDto) error {
    r.packageQueryDTO = packageQueryDTO
    r.Set("package_query_d_t_o", packageQueryDTO)
    return nil
}

func (r AlibabaIfpFulfillWarehouseTokenQueryRequest) GetPackageQueryDTO() *PackageQueryDto {
    return r.packageQueryDTO
}

