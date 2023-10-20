package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvlotteryIsvdrawAPIResponse 天猫奖池鉴权接口 API返回值
// alibaba.interact.isvlottery.isvdraw
//
// 鉴权接口，为tida.isvdraw接口鉴权
type AlibabaInteractIsvlotteryIsvdrawAPIResponse struct {
	model.CommonResponse
	AlibabaInteractIsvlotteryIsvdrawAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractIsvlotteryIsvdrawAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractIsvlotteryIsvdrawAPIResponseModel).Reset()
}

// AlibabaInteractIsvlotteryIsvdrawAPIResponseModel is 天猫奖池鉴权接口 成功返回结果
type AlibabaInteractIsvlotteryIsvdrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_isvlottery_isvdraw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无用参数
	Stub string `json:"stub,omitempty" xml:"stub,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractIsvlotteryIsvdrawAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Stub = ""
}

var poolAlibabaInteractIsvlotteryIsvdrawAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractIsvlotteryIsvdrawAPIResponse)
	},
}

// GetAlibabaInteractIsvlotteryIsvdrawAPIResponse 从 sync.Pool 获取 AlibabaInteractIsvlotteryIsvdrawAPIResponse
func GetAlibabaInteractIsvlotteryIsvdrawAPIResponse() *AlibabaInteractIsvlotteryIsvdrawAPIResponse {
	return poolAlibabaInteractIsvlotteryIsvdrawAPIResponse.Get().(*AlibabaInteractIsvlotteryIsvdrawAPIResponse)
}

// ReleaseAlibabaInteractIsvlotteryIsvdrawAPIResponse 将 AlibabaInteractIsvlotteryIsvdrawAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractIsvlotteryIsvdrawAPIResponse(v *AlibabaInteractIsvlotteryIsvdrawAPIResponse) {
	v.Reset()
	poolAlibabaInteractIsvlotteryIsvdrawAPIResponse.Put(v)
}
