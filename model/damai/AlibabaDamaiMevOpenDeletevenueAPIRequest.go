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
type AlibabaDamaiMevOpenDeletevenueAPIRequest struct {
    model.Params
    // 入参deleteVenueParam
    _deleteVenueParam   *VenueIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletevenueAPIRequest对象
func NewAlibabaDamaiMevOpenDeletevenueRequest() *AlibabaDamaiMevOpenDeletevenueAPIRequest{
    return &AlibabaDamaiMevOpenDeletevenueAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletevenue"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteVenueParam Setter
// 入参deleteVenueParam
func (r *AlibabaDamaiMevOpenDeletevenueAPIRequest) SetDeleteVenueParam(_deleteVenueParam *VenueIdOpenParam) error {
    r._deleteVenueParam = _deleteVenueParam
    r.Set("delete_venue_param", _deleteVenueParam)
    return nil
}

// DeleteVenueParam Getter
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetDeleteVenueParam() *VenueIdOpenParam {
    return r._deleteVenueParam
}
