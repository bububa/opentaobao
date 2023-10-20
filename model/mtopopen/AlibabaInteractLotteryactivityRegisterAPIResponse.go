package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractLotteryactivityRegisterAPIResponse 回传抽奖相关参数 API返回值
// alibaba.interact.lotteryactivity.register
//
// 提供接口供三方应用将数据回传到平台
type AlibabaInteractLotteryactivityRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaInteractLotteryactivityRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractLotteryactivityRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractLotteryactivityRegisterAPIResponseModel).Reset()
}

// AlibabaInteractLotteryactivityRegisterAPIResponseModel is 回传抽奖相关参数 成功返回结果
type AlibabaInteractLotteryactivityRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_lotteryactivity_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaInteractLotteryactivityRegisterResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractLotteryactivityRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaInteractLotteryactivityRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractLotteryactivityRegisterAPIResponse)
	},
}

// GetAlibabaInteractLotteryactivityRegisterAPIResponse 从 sync.Pool 获取 AlibabaInteractLotteryactivityRegisterAPIResponse
func GetAlibabaInteractLotteryactivityRegisterAPIResponse() *AlibabaInteractLotteryactivityRegisterAPIResponse {
	return poolAlibabaInteractLotteryactivityRegisterAPIResponse.Get().(*AlibabaInteractLotteryactivityRegisterAPIResponse)
}

// ReleaseAlibabaInteractLotteryactivityRegisterAPIResponse 将 AlibabaInteractLotteryactivityRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractLotteryactivityRegisterAPIResponse(v *AlibabaInteractLotteryactivityRegisterAPIResponse) {
	v.Reset()
	poolAlibabaInteractLotteryactivityRegisterAPIResponse.Put(v)
}
