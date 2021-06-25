package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
猫超商户订单拉取 APIRequest
alibaba.wdkorder.sharestock.order.get

商户拉取猫超订单数据
*/
type AlibabaWdkorderSharestockOrderGetRequest struct {
    model.Params

    // 淘宝主订单ID
    tbOrderId   int64 

}

func NewAlibabaWdkorderSharestockOrderGetRequest() *AlibabaWdkorderSharestockOrderGetRequest{
    return &AlibabaWdkorderSharestockOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkorderSharestockOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.order.get"
}

func (r AlibabaWdkorderSharestockOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkorderSharestockOrderGetRequest) SetTbOrderId(tbOrderId int64) error {
    r.tbOrderId = tbOrderId
    r.Set("tb_order_id", tbOrderId)
    return nil
}

func (r AlibabaWdkorderSharestockOrderGetRequest) GetTbOrderId() int64 {
    return r.tbOrderId
}

