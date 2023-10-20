package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductScoreGetAPIResponse 产品质量分查询 API返回值
// alibaba.icbu.product.score.get
//
// 产品质量分查询
type AlibabaIcbuProductScoreGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductScoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductScoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductScoreGetAPIResponseModel).Reset()
}

// AlibabaIcbuProductScoreGetAPIResponseModel is 产品质量分查询 成功返回结果
type AlibabaIcbuProductScoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_score_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *ProductScoreInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductScoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIcbuProductScoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductScoreGetAPIResponse)
	},
}

// GetAlibabaIcbuProductScoreGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductScoreGetAPIResponse
func GetAlibabaIcbuProductScoreGetAPIResponse() *AlibabaIcbuProductScoreGetAPIResponse {
	return poolAlibabaIcbuProductScoreGetAPIResponse.Get().(*AlibabaIcbuProductScoreGetAPIResponse)
}

// ReleaseAlibabaIcbuProductScoreGetAPIResponse 将 AlibabaIcbuProductScoreGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductScoreGetAPIResponse(v *AlibabaIcbuProductScoreGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductScoreGetAPIResponse.Put(v)
}
