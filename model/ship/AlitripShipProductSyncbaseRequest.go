package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础信息修改回调 APIRequest
alitrip.ship.product.syncbase

基础信息修改回调
*/
type AlitripShipProductSyncbaseRequest struct {
    model.Params

}

func NewAlitripShipProductSyncbaseRequest() *AlitripShipProductSyncbaseRequest{
    return &AlitripShipProductSyncbaseRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripShipProductSyncbaseRequest) GetApiMethodName() string {
    return "alitrip.ship.product.syncbase"
}

func (r AlitripShipProductSyncbaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


