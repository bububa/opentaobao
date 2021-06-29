package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼卖家主动关闭订单 API请求
alibaba.idle.isv.order.close

供外部服务商 isv 提供卖家主动关闭交易订单的功能
*/
type AlibabaIdleIsvOrderCloseRequest struct {
    model.Params
    // 输入参数
    _isvAppraiseIsvOrderCloseDto   *AppraiseIsvOrderCloseDto
}

// 初始化AlibabaIdleIsvOrderCloseRequest对象
func NewAlibabaIdleIsvOrderCloseRequest() *AlibabaIdleIsvOrderCloseRequest{
    return &AlibabaIdleIsvOrderCloseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderCloseRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.close"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvAppraiseIsvOrderCloseDto Setter
// 输入参数
func (r *AlibabaIdleIsvOrderCloseRequest) SetIsvAppraiseIsvOrderCloseDto(_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto) error {
    r._isvAppraiseIsvOrderCloseDto = _isvAppraiseIsvOrderCloseDto
    r.Set("isv_appraise_isv_order_close_dto", _isvAppraiseIsvOrderCloseDto)
    return nil
}

// IsvAppraiseIsvOrderCloseDto Getter
func (r AlibabaIdleIsvOrderCloseRequest) GetIsvAppraiseIsvOrderCloseDto() *AppraiseIsvOrderCloseDto {
    return r._isvAppraiseIsvOrderCloseDto
}
