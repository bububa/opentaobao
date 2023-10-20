package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpostradecreateAPIResponse 轻pos品牌营销下单接口 API返回值
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
type AlibabawdkpostradecreateAPIResponse struct {
	model.CommonResponse
	AlibabawdkpostradecreateAPIResponseModel
}

// AlibabawdkpostradecreateAPIResponseModel is 轻pos品牌营销下单接口 成功返回结果
type AlibabawdkpostradecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_trade_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创单结果
	Result *FastBuyPosCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
