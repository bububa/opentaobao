package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订单查询 API请求
cainiao.waybill.privacy.seller.order.get

商家查询最近100天隐私面单记录
*/
type CainiaoWaybillPrivacySellerOrderGetRequest struct {
    model.Params
}

// 初始化CainiaoWaybillPrivacySellerOrderGetRequest对象
func NewCainiaoWaybillPrivacySellerOrderGetRequest() *CainiaoWaybillPrivacySellerOrderGetRequest{
    return &CainiaoWaybillPrivacySellerOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillPrivacySellerOrderGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.privacy.seller.order.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillPrivacySellerOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
