package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保消息获取 API请求
alibaba.wdkorder.sharestock.insurance.getorder

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceGetorderRequest struct {
    model.Params
    // 淘宝子订单ID
    tbSubOrderId   int64
}

// 初始化AlibabaWdkorderSharestockInsuranceGetorderRequest对象
func NewAlibabaWdkorderSharestockInsuranceGetorderRequest() *AlibabaWdkorderSharestockInsuranceGetorderRequest{
    return &AlibabaWdkorderSharestockInsuranceGetorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceGetorderRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.insurance.getorder"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceGetorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbSubOrderId Setter
// 淘宝子订单ID
func (r *AlibabaWdkorderSharestockInsuranceGetorderRequest) SetTbSubOrderId(tbSubOrderId int64) error {
    r.tbSubOrderId = tbSubOrderId
    r.Set("tb_sub_order_id", tbSubOrderId)
    return nil
}

// TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceGetorderRequest) GetTbSubOrderId() int64 {
    return r.tbSubOrderId
}
