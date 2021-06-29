package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼卖家主动关闭订单 APIRequest
alibaba.idle.isv.order.close

供外部服务商 isv 提供卖家主动关闭交易订单的功能
*/
type AlibabaIdleIsvOrderCloseRequest struct {
    model.Params

    // 输入参数
    isvAppraiseIsvOrderCloseDto   *AppraiseIsvOrderCloseDto 

}

func NewAlibabaIdleIsvOrderCloseRequest() *AlibabaIdleIsvOrderCloseRequest{
    return &AlibabaIdleIsvOrderCloseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvOrderCloseRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.close"
}

func (r AlibabaIdleIsvOrderCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvOrderCloseRequest) SetIsvAppraiseIsvOrderCloseDto(isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto) error {
    r.isvAppraiseIsvOrderCloseDto = isvAppraiseIsvOrderCloseDto
    r.Set("isv_appraise_isv_order_close_dto", isvAppraiseIsvOrderCloseDto)
    return nil
}

func (r AlibabaIdleIsvOrderCloseRequest) GetIsvAppraiseIsvOrderCloseDto() *AppraiseIsvOrderCloseDto {
    return r.isvAppraiseIsvOrderCloseDto
}

