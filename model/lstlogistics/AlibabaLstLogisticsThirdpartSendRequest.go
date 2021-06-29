package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-使用第三方物流发货 APIRequest
alibaba.lst.logistics.thirdpart.send

异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsThirdpartSendRequest struct {
    model.Params

    // 入参
    param   *SendOfflineOrderParam 

}

func NewAlibabaLstLogisticsThirdpartSendRequest() *AlibabaLstLogisticsThirdpartSendRequest{
    return &AlibabaLstLogisticsThirdpartSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstLogisticsThirdpartSendRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.thirdpart.send"
}

func (r AlibabaLstLogisticsThirdpartSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstLogisticsThirdpartSendRequest) SetParam(param *SendOfflineOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstLogisticsThirdpartSendRequest) GetParam() *SendOfflineOrderParam {
    return r.param
}

