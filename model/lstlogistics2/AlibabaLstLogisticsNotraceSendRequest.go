package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-无需物流发货 APIRequest
alibaba.lst.logistics.notrace.send

异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsNotraceSendRequest struct {
    model.Params

    // 入参
    param   *SendDummyOrderParam 

}

func NewAlibabaLstLogisticsNotraceSendRequest() *AlibabaLstLogisticsNotraceSendRequest{
    return &AlibabaLstLogisticsNotraceSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstLogisticsNotraceSendRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.notrace.send"
}

func (r AlibabaLstLogisticsNotraceSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstLogisticsNotraceSendRequest) SetParam(param *SendDummyOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstLogisticsNotraceSendRequest) GetParam() *SendDummyOrderParam {
    return r.param
}

