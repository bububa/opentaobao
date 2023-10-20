package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderScoreAPIResponse 订单打分和评价 API返回值
// alibaba.happytrip.taxi.order.score
//
// 对司机进行评分，只有订单结束后，才能进行。
type AlibabaHappytripTaxiOrderScoreAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderScoreAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderScoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderScoreAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderScoreAPIResponseModel is 订单打分和评价 成功返回结果
type AlibabaHappytripTaxiOrderScoreAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_score_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderScoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
}

var poolAlibabaHappytripTaxiOrderScoreAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderScoreAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderScoreAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderScoreAPIResponse
func GetAlibabaHappytripTaxiOrderScoreAPIResponse() *AlibabaHappytripTaxiOrderScoreAPIResponse {
	return poolAlibabaHappytripTaxiOrderScoreAPIResponse.Get().(*AlibabaHappytripTaxiOrderScoreAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderScoreAPIResponse 将 AlibabaHappytripTaxiOrderScoreAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderScoreAPIResponse(v *AlibabaHappytripTaxiOrderScoreAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderScoreAPIResponse.Put(v)
}
