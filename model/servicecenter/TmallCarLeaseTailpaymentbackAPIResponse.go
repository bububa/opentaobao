package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
尾款处置方案回传 API返回值 
tmall.car.lease.tailpaymentback

尾款处置方案回传
*/
type TmallCarLeaseTailpaymentbackAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseTailpaymentbackAPIResponseModel
}

// 尾款处置方案回传 成功返回结果
type TmallCarLeaseTailpaymentbackAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_car_lease_tailpaymentback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallCarLeaseTailpaymentbackResult `json:"result,omitempty" xml:"result,omitempty"`
}
