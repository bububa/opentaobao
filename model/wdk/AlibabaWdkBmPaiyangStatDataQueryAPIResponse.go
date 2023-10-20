package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmPaiyangStatDataQueryAPIResponse 派样统计数据查询 API返回值
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
type AlibabaWdkBmPaiyangStatDataQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBmPaiyangStatDataQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkBmPaiyangStatDataQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkBmPaiyangStatDataQueryAPIResponseModel).Reset()
}

// AlibabaWdkBmPaiyangStatDataQueryAPIResponseModel is 派样统计数据查询 成功返回结果
type AlibabaWdkBmPaiyangStatDataQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_paiyang_stat_data_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *BmPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkBmPaiyangStatDataQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkBmPaiyangStatDataQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBmPaiyangStatDataQueryAPIResponse)
	},
}

// GetAlibabaWdkBmPaiyangStatDataQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkBmPaiyangStatDataQueryAPIResponse
func GetAlibabaWdkBmPaiyangStatDataQueryAPIResponse() *AlibabaWdkBmPaiyangStatDataQueryAPIResponse {
	return poolAlibabaWdkBmPaiyangStatDataQueryAPIResponse.Get().(*AlibabaWdkBmPaiyangStatDataQueryAPIResponse)
}

// ReleaseAlibabaWdkBmPaiyangStatDataQueryAPIResponse 将 AlibabaWdkBmPaiyangStatDataQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkBmPaiyangStatDataQueryAPIResponse(v *AlibabaWdkBmPaiyangStatDataQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkBmPaiyangStatDataQueryAPIResponse.Put(v)
}
