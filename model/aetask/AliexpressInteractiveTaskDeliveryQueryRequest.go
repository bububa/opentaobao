package aetask

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE互动任务投放 API请求
aliexpress.interactive.task.delivery.query

将内部配置好的任务，如浏览商品，店铺投放给外部ISV
*/
type AliexpressInteractiveTaskDeliveryQueryRequest struct {
    model.Params
    // 返回结果
    requestDto   *QueryDeliveryRequestDto
}

// 初始化AliexpressInteractiveTaskDeliveryQueryRequest对象
func NewAliexpressInteractiveTaskDeliveryQueryRequest() *AliexpressInteractiveTaskDeliveryQueryRequest{
    return &AliexpressInteractiveTaskDeliveryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressInteractiveTaskDeliveryQueryRequest) GetApiMethodName() string {
    return "aliexpress.interactive.task.delivery.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressInteractiveTaskDeliveryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestDto Setter
// 返回结果
func (r *AliexpressInteractiveTaskDeliveryQueryRequest) SetRequestDto(requestDto *QueryDeliveryRequestDto) error {
    r.requestDto = requestDto
    r.Set("request_dto", requestDto)
    return nil
}

// RequestDto Getter
func (r AliexpressInteractiveTaskDeliveryQueryRequest) GetRequestDto() *QueryDeliveryRequestDto {
    return r.requestDto
}
