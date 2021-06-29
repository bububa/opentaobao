package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单打分和评价 API返回值 
alibaba.happytrip.taxi.order.score

对司机进行评分，只有订单结束后，才能进行。
*/
type AlibabaHappytripTaxiOrderScoreAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiOrderScoreResponse
}

// 订单打分和评价 成功返回结果
type AlibabaHappytripTaxiOrderScoreResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_score_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`
    // 错误信息
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
