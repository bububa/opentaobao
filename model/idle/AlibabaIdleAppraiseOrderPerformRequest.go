package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼验货担保服务商订单履约V1 API请求
alibaba.idle.appraise.order.perform

闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
*/
type AlibabaIdleAppraiseOrderPerformRequest struct {
    model.Params
    // AppraiseOrderSynDto
    appraiseOrderSynDto   *AppraiseOrderSynDto
}

// 初始化AlibabaIdleAppraiseOrderPerformRequest对象
func NewAlibabaIdleAppraiseOrderPerformRequest() *AlibabaIdleAppraiseOrderPerformRequest{
    return &AlibabaIdleAppraiseOrderPerformRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseOrderPerformRequest) GetApiMethodName() string {
    return "alibaba.idle.appraise.order.perform"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseOrderPerformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppraiseOrderSynDto Setter
// AppraiseOrderSynDto
func (r *AlibabaIdleAppraiseOrderPerformRequest) SetAppraiseOrderSynDto(appraiseOrderSynDto *AppraiseOrderSynDto) error {
    r.appraiseOrderSynDto = appraiseOrderSynDto
    r.Set("appraise_order_syn_dto", appraiseOrderSynDto)
    return nil
}

// AppraiseOrderSynDto Getter
func (r AlibabaIdleAppraiseOrderPerformRequest) GetAppraiseOrderSynDto() *AppraiseOrderSynDto {
    return r.appraiseOrderSynDto
}
