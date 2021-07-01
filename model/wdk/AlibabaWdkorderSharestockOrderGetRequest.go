package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
猫超商户订单拉取 API请求
alibaba.wdkorder.sharestock.order.get

商户拉取猫超订单数据
*/
type AlibabaWdkorderSharestockOrderGetAPIRequest struct {
    model.Params
    // 淘宝主订单ID
    _tbOrderId   int64
}

// 初始化AlibabaWdkorderSharestockOrderGetAPIRequest对象
func NewAlibabaWdkorderSharestockOrderGetRequest() *AlibabaWdkorderSharestockOrderGetAPIRequest{
    return &AlibabaWdkorderSharestockOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbOrderId Setter
// 淘宝主订单ID
func (r *AlibabaWdkorderSharestockOrderGetAPIRequest) SetTbOrderId(_tbOrderId int64) error {
    r._tbOrderId = _tbOrderId
    r.Set("tb_order_id", _tbOrderId)
    return nil
}

// TbOrderId Getter
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetTbOrderId() int64 {
    return r._tbOrderId
}
