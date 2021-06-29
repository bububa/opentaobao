package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场馆接口deleteVenue API请求
alibaba.damai.mev.open.deletevenue

开放接口，删除场馆
*/
type AlibabaDamaiMevOpenDeletevenueRequest struct {
    model.Params
    // 入参deleteVenueParam
    deleteVenueParam   *VenueIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletevenueRequest对象
func NewAlibabaDamaiMevOpenDeletevenueRequest() *AlibabaDamaiMevOpenDeletevenueRequest{
    return &AlibabaDamaiMevOpenDeletevenueRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletevenueRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletevenue"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletevenueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteVenueParam Setter
// 入参deleteVenueParam
func (r *AlibabaDamaiMevOpenDeletevenueRequest) SetDeleteVenueParam(deleteVenueParam *VenueIdOpenParam) error {
    r.deleteVenueParam = deleteVenueParam
    r.Set("delete_venue_param", deleteVenueParam)
    return nil
}

// DeleteVenueParam Getter
func (r AlibabaDamaiMevOpenDeletevenueRequest) GetDeleteVenueParam() *VenueIdOpenParam {
    return r.deleteVenueParam
}
