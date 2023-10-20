package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardordersStatusSendedAPIResponse 制卡商通知实体卡发货完成 API返回值
// alibaba.fundplatform.cardorders.status.sended
//
// 当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
type AlibabaFundplatformCardordersStatusSendedAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardordersStatusSendedAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardordersStatusSendedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformCardordersStatusSendedAPIResponseModel).Reset()
}

// AlibabaFundplatformCardordersStatusSendedAPIResponseModel is 制卡商通知实体卡发货完成 成功返回结果
type AlibabaFundplatformCardordersStatusSendedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_status_sended_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CardMakingInformResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardordersStatusSendedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFundplatformCardordersStatusSendedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformCardordersStatusSendedAPIResponse)
	},
}

// GetAlibabaFundplatformCardordersStatusSendedAPIResponse 从 sync.Pool 获取 AlibabaFundplatformCardordersStatusSendedAPIResponse
func GetAlibabaFundplatformCardordersStatusSendedAPIResponse() *AlibabaFundplatformCardordersStatusSendedAPIResponse {
	return poolAlibabaFundplatformCardordersStatusSendedAPIResponse.Get().(*AlibabaFundplatformCardordersStatusSendedAPIResponse)
}

// ReleaseAlibabaFundplatformCardordersStatusSendedAPIResponse 将 AlibabaFundplatformCardordersStatusSendedAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformCardordersStatusSendedAPIResponse(v *AlibabaFundplatformCardordersStatusSendedAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformCardordersStatusSendedAPIResponse.Put(v)
}
