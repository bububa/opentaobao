package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXianyuTenderOrderPerformAPIResponse 闲鱼暗拍订单履约 API返回值
// alibaba.xianyu.tender.order.perform
//
// 闲鱼暗拍订单履约
type AlibabaXianyuTenderOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaXianyuTenderOrderPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaXianyuTenderOrderPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaXianyuTenderOrderPerformAPIResponseModel).Reset()
}

// AlibabaXianyuTenderOrderPerformAPIResponseModel is 闲鱼暗拍订单履约 成功返回结果
type AlibabaXianyuTenderOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xianyu_tender_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaXianyuTenderOrderPerformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaXianyuTenderOrderPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaXianyuTenderOrderPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaXianyuTenderOrderPerformAPIResponse)
	},
}

// GetAlibabaXianyuTenderOrderPerformAPIResponse 从 sync.Pool 获取 AlibabaXianyuTenderOrderPerformAPIResponse
func GetAlibabaXianyuTenderOrderPerformAPIResponse() *AlibabaXianyuTenderOrderPerformAPIResponse {
	return poolAlibabaXianyuTenderOrderPerformAPIResponse.Get().(*AlibabaXianyuTenderOrderPerformAPIResponse)
}

// ReleaseAlibabaXianyuTenderOrderPerformAPIResponse 将 AlibabaXianyuTenderOrderPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaXianyuTenderOrderPerformAPIResponse(v *AlibabaXianyuTenderOrderPerformAPIResponse) {
	v.Reset()
	poolAlibabaXianyuTenderOrderPerformAPIResponse.Put(v)
}
