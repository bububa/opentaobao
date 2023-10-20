package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderComplaintGetAPIResponse 投诉详情 API返回值
// alibaba.happytrip.taxi.order.complaint.get
//
// 获取投诉处理进度详情
type AlibabaHappytripTaxiOrderComplaintGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderComplaintGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderComplaintGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderComplaintGetAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderComplaintGetAPIResponseModel is 投诉详情 成功返回结果
type AlibabaHappytripTaxiOrderComplaintGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_complaint_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 投诉详情获取结果
	Data *AlibabaHappytripTaxiOrderComplaintGetStruct `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderComplaintGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiOrderComplaintGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderComplaintGetAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderComplaintGetAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderComplaintGetAPIResponse
func GetAlibabaHappytripTaxiOrderComplaintGetAPIResponse() *AlibabaHappytripTaxiOrderComplaintGetAPIResponse {
	return poolAlibabaHappytripTaxiOrderComplaintGetAPIResponse.Get().(*AlibabaHappytripTaxiOrderComplaintGetAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderComplaintGetAPIResponse 将 AlibabaHappytripTaxiOrderComplaintGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderComplaintGetAPIResponse(v *AlibabaHappytripTaxiOrderComplaintGetAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderComplaintGetAPIResponse.Put(v)
}
