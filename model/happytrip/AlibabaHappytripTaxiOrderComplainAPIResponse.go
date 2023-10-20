package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderComplainAPIResponse 用户投诉 API返回值
// alibaba.happytrip.taxi.order.complain
//
// 一个订单只能投诉一次，不可重复投诉
//
// 投诉选项
// 0	其他原因；
// 1	司机原因导致行程被取消；
// 2	服务态度恶劣；
// 3	未坐车产生费用；
// 4	额外收取不合理费用；
// 5	路不熟多产生费用；
// 6	提前计费；
// 7	未及时结束计费；
// 8	司机绕路；
// 9	司机迟到；
// 10	司机爽约或拒载；
// 11	骚扰乘客；
// 12	危险驾驶；
// 13	不是订单显示车辆或司机；
type AlibabaHappytripTaxiOrderComplainAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderComplainAPIResponseModel
}

// AlibabaHappytripTaxiOrderComplainAPIResponseModel is 用户投诉 成功返回结果
type AlibabaHappytripTaxiOrderComplainAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_complain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 投诉结果
	Data *AlibabaHappytripTaxiOrderComplainStruct `json:"data,omitempty" xml:"data,omitempty"`
}
