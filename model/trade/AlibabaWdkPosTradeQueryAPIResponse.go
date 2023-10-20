package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeQueryAPIResponse 轻pos品牌营销查询接口 API返回值
// alibaba.wdk.pos.trade.query
//
// 轻pos品牌营销场景，外部商家查询营销信息
type AlibabaWdkPosTradeQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPosTradeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPosTradeQueryAPIResponseModel).Reset()
}

// AlibabaWdkPosTradeQueryAPIResponseModel is 轻pos品牌营销查询接口 成功返回结果
type AlibabaWdkPosTradeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_trade_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *FastBuyPosQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPosTradeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPosTradeQueryAPIResponse)
	},
}

// GetAlibabaWdkPosTradeQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkPosTradeQueryAPIResponse
func GetAlibabaWdkPosTradeQueryAPIResponse() *AlibabaWdkPosTradeQueryAPIResponse {
	return poolAlibabaWdkPosTradeQueryAPIResponse.Get().(*AlibabaWdkPosTradeQueryAPIResponse)
}

// ReleaseAlibabaWdkPosTradeQueryAPIResponse 将 AlibabaWdkPosTradeQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPosTradeQueryAPIResponse(v *AlibabaWdkPosTradeQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkPosTradeQueryAPIResponse.Put(v)
}
