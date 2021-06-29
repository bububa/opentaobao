package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收订单履约V2 APIRequest
alibaba.idle.recycle.order.perform

闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
*/
type AlibabaIdleRecycleOrderPerformRequest struct {
    model.Params

    // 参数
    param0   *RecycleOrderSynDto 

}

func NewAlibabaIdleRecycleOrderPerformRequest() *AlibabaIdleRecycleOrderPerformRequest{
    return &AlibabaIdleRecycleOrderPerformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRecycleOrderPerformRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.perform"
}

func (r AlibabaIdleRecycleOrderPerformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRecycleOrderPerformRequest) SetParam0(param0 *RecycleOrderSynDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaIdleRecycleOrderPerformRequest) GetParam0() *RecycleOrderSynDto {
    return r.param0
}

