package ascpqcc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpQccSampleCancelItemRelationAPIResponse 魅力惠样品解除父子商品关系 API返回值
// alibaba.ascp.qcc.sample.cancel.item.relation
//
// 品控中心魅力惠样品解除父子商品关系
type AlibabaAscpQccSampleCancelItemRelationAPIResponse struct {
	model.CommonResponse
	AlibabaAscpQccSampleCancelItemRelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpQccSampleCancelItemRelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpQccSampleCancelItemRelationAPIResponseModel).Reset()
}

// AlibabaAscpQccSampleCancelItemRelationAPIResponseModel is 魅力惠样品解除父子商品关系 成功返回结果
type AlibabaAscpQccSampleCancelItemRelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_qcc_sample_cancel_item_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpQccSampleCancelItemRelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpQccSampleCancelItemRelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpQccSampleCancelItemRelationAPIResponse)
	},
}

// GetAlibabaAscpQccSampleCancelItemRelationAPIResponse 从 sync.Pool 获取 AlibabaAscpQccSampleCancelItemRelationAPIResponse
func GetAlibabaAscpQccSampleCancelItemRelationAPIResponse() *AlibabaAscpQccSampleCancelItemRelationAPIResponse {
	return poolAlibabaAscpQccSampleCancelItemRelationAPIResponse.Get().(*AlibabaAscpQccSampleCancelItemRelationAPIResponse)
}

// ReleaseAlibabaAscpQccSampleCancelItemRelationAPIResponse 将 AlibabaAscpQccSampleCancelItemRelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpQccSampleCancelItemRelationAPIResponse(v *AlibabaAscpQccSampleCancelItemRelationAPIResponse) {
	v.Reset()
	poolAlibabaAscpQccSampleCancelItemRelationAPIResponse.Put(v)
}
