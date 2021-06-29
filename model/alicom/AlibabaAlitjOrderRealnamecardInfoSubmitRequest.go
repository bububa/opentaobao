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
    _orderNo   int64
    // sim卡iccid（一般为18位到20位）
    _iccid   string
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
func (r *AlibabaAlitjOrderRealnamecardInfoSubmitRequest) SetOrderNo(_orderNo int64) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetOrderNo() int64 {
    return r._orderNo
}
// Iccid Setter
// sim卡iccid（一般为18位到20位）
func (r *AlibabaAlitjOrderRealnamecardInfoSubmitRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetIccid() string {
    return r._iccid
}
