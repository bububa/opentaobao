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
type AlibabaWdkorderSharestockOrderGetRequest struct {
    model.Params
    // 淘宝主订单ID
    tbOrderId   int64
}

// 初始化AlibabaWdkorderSharestockOrderGetRequest对象
func NewAlibabaWdkorderSharestockOrderGetRequest() *AlibabaWdkorderSharestockOrderGetRequest{
    return &AlibabaWdkorderSharestockOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbOrderId Setter
// 淘宝主订单ID
func (r *AlibabaWdkorderSharestockOrderGetRequest) SetTbOrderId(tbOrderId int64) error {
    r.tbOrderId = tbOrderId
    r.Set("tb_order_id", tbOrderId)
    return nil
}

// TbOrderId Getter
func (r AlibabaWdkorderSharestockOrderGetRequest) GetTbOrderId() int64 {
    return r.tbOrderId
}
