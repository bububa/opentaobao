package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
BMS出库通知 API请求
cainiao.bms.order.consign.confirm

BMS出库后，通知ISV
*/
type CainiaoBmsOrderConsignConfirmRequest struct {
    model.Params
    // 通知消息主体
    content   *BmsConsignOrderConfirm
}

// 初始化CainiaoBmsOrderConsignConfirmRequest对象
func NewCainiaoBmsOrderConsignConfirmRequest() *CainiaoBmsOrderConsignConfirmRequest{
    return &CainiaoBmsOrderConsignConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoBmsOrderConsignConfirmRequest) GetApiMethodName() string {
    return "cainiao.bms.order.consign.confirm"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoBmsOrderConsignConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 通知消息主体
func (r *CainiaoBmsOrderConsignConfirmRequest) SetContent(content *BmsConsignOrderConfirm) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r CainiaoBmsOrderConsignConfirmRequest) GetContent() *BmsConsignOrderConfirm {
    return r.content
}
