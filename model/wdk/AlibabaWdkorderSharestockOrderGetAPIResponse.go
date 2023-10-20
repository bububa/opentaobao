package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockOrderGetAPIResponse 猫超商户订单拉取 API返回值
// alibaba.wdkorder.sharestock.order.get
//
// 商户拉取猫超订单数据
type AlibabaWdkorderSharestockOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkorderSharestockOrderGetAPIResponseModel).Reset()
}

// AlibabaWdkorderSharestockOrderGetAPIResponseModel is 猫超商户订单拉取 成功返回结果
type AlibabaWdkorderSharestockOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *MaochaoOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkorderSharestockOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockOrderGetAPIResponse)
	},
}

// GetAlibabaWdkorderSharestockOrderGetAPIResponse 从 sync.Pool 获取 AlibabaWdkorderSharestockOrderGetAPIResponse
func GetAlibabaWdkorderSharestockOrderGetAPIResponse() *AlibabaWdkorderSharestockOrderGetAPIResponse {
	return poolAlibabaWdkorderSharestockOrderGetAPIResponse.Get().(*AlibabaWdkorderSharestockOrderGetAPIResponse)
}

// ReleaseAlibabaWdkorderSharestockOrderGetAPIResponse 将 AlibabaWdkorderSharestockOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkorderSharestockOrderGetAPIResponse(v *AlibabaWdkorderSharestockOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkorderSharestockOrderGetAPIResponse.Put(v)
}
