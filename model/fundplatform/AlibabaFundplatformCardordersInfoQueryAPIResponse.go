package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardordersInfoQueryAPIResponse 根据制卡单分页查询卡信息 API返回值
// alibaba.fundplatform.cardorders.info.query
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
type AlibabaFundplatformCardordersInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardordersInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardordersInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformCardordersInfoQueryAPIResponseModel).Reset()
}

// AlibabaFundplatformCardordersInfoQueryAPIResponseModel is 根据制卡单分页查询卡信息 成功返回结果
type AlibabaFundplatformCardordersInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CardMakingInfoQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardordersInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFundplatformCardordersInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformCardordersInfoQueryAPIResponse)
	},
}

// GetAlibabaFundplatformCardordersInfoQueryAPIResponse 从 sync.Pool 获取 AlibabaFundplatformCardordersInfoQueryAPIResponse
func GetAlibabaFundplatformCardordersInfoQueryAPIResponse() *AlibabaFundplatformCardordersInfoQueryAPIResponse {
	return poolAlibabaFundplatformCardordersInfoQueryAPIResponse.Get().(*AlibabaFundplatformCardordersInfoQueryAPIResponse)
}

// ReleaseAlibabaFundplatformCardordersInfoQueryAPIResponse 将 AlibabaFundplatformCardordersInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformCardordersInfoQueryAPIResponse(v *AlibabaFundplatformCardordersInfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformCardordersInfoQueryAPIResponse.Put(v)
}
