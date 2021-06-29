package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-无需物流发货 API请求
alibaba.lst.logistics.notrace.send

异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsNotraceSendRequest struct {
    model.Params
    // 入参
    param   *SendDummyOrderParam
}

// 初始化AlibabaLstLogisticsNotraceSendRequest对象
func NewAlibabaLstLogisticsNotraceSendRequest() *AlibabaLstLogisticsNotraceSendRequest{
    return &AlibabaLstLogisticsNotraceSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsNotraceSendRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.notrace.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsNotraceSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaLstLogisticsNotraceSendRequest) SetParam(param *SendDummyOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaLstLogisticsNotraceSendRequest) GetParam() *SendDummyOrderParam {
    return r.param
}
