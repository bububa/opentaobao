package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsOrderGetAPIResponse 查询店仓作业单据清单 （库存对账辅助）-回流单 API返回值
// alibaba.wdk.ums.order.get
//
// 查询店仓作业单据清单 （库存对账辅助）-回流单
type AlibabaWdkUmsOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkUmsOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkUmsOrderGetAPIResponseModel).Reset()
}

// AlibabaWdkUmsOrderGetAPIResponseModel is 查询店仓作业单据清单 （库存对账辅助）-回流单 成功返回结果
type AlibabaWdkUmsOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkUmsOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkUmsOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsOrderGetAPIResponse)
	},
}

// GetAlibabaWdkUmsOrderGetAPIResponse 从 sync.Pool 获取 AlibabaWdkUmsOrderGetAPIResponse
func GetAlibabaWdkUmsOrderGetAPIResponse() *AlibabaWdkUmsOrderGetAPIResponse {
	return poolAlibabaWdkUmsOrderGetAPIResponse.Get().(*AlibabaWdkUmsOrderGetAPIResponse)
}

// ReleaseAlibabaWdkUmsOrderGetAPIResponse 将 AlibabaWdkUmsOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkUmsOrderGetAPIResponse(v *AlibabaWdkUmsOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkUmsOrderGetAPIResponse.Put(v)
}
