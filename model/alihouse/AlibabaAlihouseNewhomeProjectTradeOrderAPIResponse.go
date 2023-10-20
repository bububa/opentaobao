package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse 旺铺楼盘和交易商品排序 API返回值
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
type AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectTradeOrderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectTradeOrderAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectTradeOrderAPIResponseModel is 旺铺楼盘和交易商品排序 成功返回结果
type AlibabaAlihouseNewhomeProjectTradeOrderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_trade_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectTradeOrderResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectTradeOrderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse
func GetAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse() *AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse 将 AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse(v *AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectTradeOrderAPIResponse.Put(v)
}
