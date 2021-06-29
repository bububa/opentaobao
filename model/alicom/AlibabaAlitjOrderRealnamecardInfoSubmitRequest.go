package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里实人认证卡片信息回传 API请求
alibaba.alitj.order.realnamecard.info.submit

阿里实人认证卡片信息回传。ISP相关商家在线对接阿里通信的实人认证功能，在线提交订单对应运营商的合约订购相关信息，以便完成在线使用实人认证功能。
*/
type AlibabaAlitjOrderRealnamecardInfoSubmitRequest struct {
    model.Params
    // 淘宝订单号
    orderNo   int64
    // sim卡iccid（一般为18位到20位）
    iccid   string
}

// 初始化AlibabaAlitjOrderRealnamecardInfoSubmitRequest对象
func NewAlibabaAlitjOrderRealnamecardInfoSubmitRequest() *AlibabaAlitjOrderRealnamecardInfoSubmitRequest{
    return &AlibabaAlitjOrderRealnamecardInfoSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetApiMethodName() string {
    return "alibaba.alitj.order.realnamecard.info.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 淘宝订单号
func (r *AlibabaAlitjOrderRealnamecardInfoSubmitRequest) SetOrderNo(orderNo int64) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetOrderNo() int64 {
    return r.orderNo
}
// Iccid Setter
// sim卡iccid（一般为18位到20位）
func (r *AlibabaAlitjOrderRealnamecardInfoSubmitRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetIccid() string {
    return r.iccid
}
