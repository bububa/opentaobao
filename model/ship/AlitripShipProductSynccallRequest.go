package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全量同步回调 APIRequest
alitrip.ship.product.synccall

全量同步接口
*/
type AlitripShipProductSynccallRequest struct {
    model.Params

}

func NewAlitripShipProductSynccallRequest() *AlitripShipProductSynccallRequest{
    return &AlitripShipProductSynccallRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripShipProductSynccallRequest) GetApiMethodName() string {
    return "alitrip.ship.product.synccall"
}

func (r AlitripShipProductSynccallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


