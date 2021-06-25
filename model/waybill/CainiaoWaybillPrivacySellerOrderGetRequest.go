package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订单查询 APIRequest
cainiao.waybill.privacy.seller.order.get

商家查询最近100天隐私面单记录
*/
type CainiaoWaybillPrivacySellerOrderGetRequest struct {
    model.Params

}

func NewCainiaoWaybillPrivacySellerOrderGetRequest() *CainiaoWaybillPrivacySellerOrderGetRequest{
    return &CainiaoWaybillPrivacySellerOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillPrivacySellerOrderGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.privacy.seller.order.get"
}

func (r CainiaoWaybillPrivacySellerOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


