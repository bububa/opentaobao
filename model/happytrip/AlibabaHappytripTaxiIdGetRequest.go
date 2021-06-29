package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取请求id APIRequest
alibaba.happytrip.taxi.id.get

获取订单号
*/
type AlibabaHappytripTaxiIdGetRequest struct {
    model.Params

    // 用户唯一标识
    uid   string 

}

func NewAlibabaHappytripTaxiIdGetRequest() *AlibabaHappytripTaxiIdGetRequest{
    return &AlibabaHappytripTaxiIdGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiIdGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.id.get"
}

func (r AlibabaHappytripTaxiIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiIdGetRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r AlibabaHappytripTaxiIdGetRequest) GetUid() string {
    return r.uid
}

