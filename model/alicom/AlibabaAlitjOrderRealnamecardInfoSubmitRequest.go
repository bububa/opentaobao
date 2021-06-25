package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里实人认证卡片信息回传 APIRequest
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

func NewAlibabaAlitjOrderRealnamecardInfoSubmitRequest() *AlibabaAlitjOrderRealnamecardInfoSubmitRequest{
    return &AlibabaAlitjOrderRealnamecardInfoSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetApiMethodName() string {
    return "alibaba.alitj.order.realnamecard.info.submit"
}

func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlitjOrderRealnamecardInfoSubmitRequest) SetOrderNo(orderNo int64) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetOrderNo() int64 {
    return r.orderNo
}

func (r *AlibabaAlitjOrderRealnamecardInfoSubmitRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

func (r AlibabaAlitjOrderRealnamecardInfoSubmitRequest) GetIccid() string {
    return r.iccid
}

