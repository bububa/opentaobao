package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmTradeActivityQueryAPIResponse 品牌营销的订单活动信息查询 API返回值
// alibaba.wdk.bm.trade.activity.query
//
// 品牌营销的订单活动信息查询
type AlibabaWdkBmTradeActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBmTradeActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkBmTradeActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkBmTradeActivityQueryAPIResponseModel).Reset()
}

// AlibabaWdkBmTradeActivityQueryAPIResponseModel is 品牌营销的订单活动信息查询 成功返回结果
type AlibabaWdkBmTradeActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_trade_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果数据
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkBmTradeActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkBmTradeActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBmTradeActivityQueryAPIResponse)
	},
}

// GetAlibabaWdkBmTradeActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkBmTradeActivityQueryAPIResponse
func GetAlibabaWdkBmTradeActivityQueryAPIResponse() *AlibabaWdkBmTradeActivityQueryAPIResponse {
	return poolAlibabaWdkBmTradeActivityQueryAPIResponse.Get().(*AlibabaWdkBmTradeActivityQueryAPIResponse)
}

// ReleaseAlibabaWdkBmTradeActivityQueryAPIResponse 将 AlibabaWdkBmTradeActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkBmTradeActivityQueryAPIResponse(v *AlibabaWdkBmTradeActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkBmTradeActivityQueryAPIResponse.Put(v)
}
