package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单指派 API返回值 
alibaba.happytrip.taxi.order.assign

通知供应商订单指派成功
*/
type AlibabaHappytripTaxiOrderAssignAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiOrderAssignResponse
}

// 订单指派 成功返回结果
type AlibabaHappytripTaxiOrderAssignResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_assign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误代码
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`
    // 错误描述
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
