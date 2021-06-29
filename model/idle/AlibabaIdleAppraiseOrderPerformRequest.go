package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼验货担保服务商订单履约V1 APIRequest
alibaba.idle.appraise.order.perform

闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
*/
type AlibabaIdleAppraiseOrderPerformRequest struct {
    model.Params

    // AppraiseOrderSynDto
    appraiseOrderSynDto   *AppraiseOrderSynDto 

}

func NewAlibabaIdleAppraiseOrderPerformRequest() *AlibabaIdleAppraiseOrderPerformRequest{
    return &AlibabaIdleAppraiseOrderPerformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleAppraiseOrderPerformRequest) GetApiMethodName() string {
    return "alibaba.idle.appraise.order.perform"
}

func (r AlibabaIdleAppraiseOrderPerformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleAppraiseOrderPerformRequest) SetAppraiseOrderSynDto(appraiseOrderSynDto *AppraiseOrderSynDto) error {
    r.appraiseOrderSynDto = appraiseOrderSynDto
    r.Set("appraise_order_syn_dto", appraiseOrderSynDto)
    return nil
}

func (r AlibabaIdleAppraiseOrderPerformRequest) GetAppraiseOrderSynDto() *AppraiseOrderSynDto {
    return r.appraiseOrderSynDto
}

