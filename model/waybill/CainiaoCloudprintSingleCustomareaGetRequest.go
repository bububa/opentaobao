package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家单一自定义区 APIRequest
cainiao.cloudprint.single.customarea.get

商家所有快递公司模板只有一个自定义区
*/
type CainiaoCloudprintSingleCustomareaGetRequest struct {
    model.Params

    // 这是商家用户id
    sellerId   int64 

}

func NewCainiaoCloudprintSingleCustomareaGetRequest() *CainiaoCloudprintSingleCustomareaGetRequest{
    return &CainiaoCloudprintSingleCustomareaGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintSingleCustomareaGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.single.customarea.get"
}

func (r CainiaoCloudprintSingleCustomareaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCloudprintSingleCustomareaGetRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r CainiaoCloudprintSingleCustomareaGetRequest) GetSellerId() int64 {
    return r.sellerId
}

