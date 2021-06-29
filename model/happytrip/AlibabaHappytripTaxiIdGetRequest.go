package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取请求id API请求
alibaba.happytrip.taxi.id.get

获取订单号
*/
type AlibabaHappytripTaxiIdGetRequest struct {
    model.Params
    // 用户唯一标识
    uid   string
}

// 初始化AlibabaHappytripTaxiIdGetRequest对象
func NewAlibabaHappytripTaxiIdGetRequest() *AlibabaHappytripTaxiIdGetRequest{
    return &AlibabaHappytripTaxiIdGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiIdGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.id.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiIdGetRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiIdGetRequest) GetUid() string {
    return r.uid
}
