package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderNotifyAPIResponse 状态通知 API返回值
// alibaba.happytrip.taxi.order.notify
//
// 当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。
type AlibabaHappytripTaxiOrderNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderNotifyAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderNotifyAPIResponseModel is 状态通知 成功返回结果
type AlibabaHappytripTaxiOrderNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，0为通知成功，非0为通知失败
	Errno string `json:"errno,omitempty" xml:"errno,omitempty"`
	// 错误信息描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errno = ""
	m.Errmsg = ""
}

var poolAlibabaHappytripTaxiOrderNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderNotifyAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderNotifyAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderNotifyAPIResponse
func GetAlibabaHappytripTaxiOrderNotifyAPIResponse() *AlibabaHappytripTaxiOrderNotifyAPIResponse {
	return poolAlibabaHappytripTaxiOrderNotifyAPIResponse.Get().(*AlibabaHappytripTaxiOrderNotifyAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderNotifyAPIResponse 将 AlibabaHappytripTaxiOrderNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderNotifyAPIResponse(v *AlibabaHappytripTaxiOrderNotifyAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderNotifyAPIResponse.Put(v)
}
