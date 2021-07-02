package interact

import (
	"encoding/xml"

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

// AlibabaInteractIsvlotteryIsvdrawAPIResponseModel is 天猫奖池鉴权接口 成功返回结果
type AlibabaInteractIsvlotteryIsvdrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_isvlottery_isvdraw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无用参数
	Stub string `json:"stub,omitempty" xml:"stub,omitempty"`
}
