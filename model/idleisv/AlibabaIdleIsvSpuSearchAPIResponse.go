package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvSpuSearchAPIResponse spu搜索接口 API返回值
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
type AlibabaIdleIsvSpuSearchAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvSpuSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvSpuSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvSpuSearchAPIResponseModel).Reset()
}

// AlibabaIdleIsvSpuSearchAPIResponseModel is spu搜索接口 成功返回结果
type AlibabaIdleIsvSpuSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_spu_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleIsvSpuSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvSpuSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvSpuSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvSpuSearchAPIResponse)
	},
}

// GetAlibabaIdleIsvSpuSearchAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvSpuSearchAPIResponse
func GetAlibabaIdleIsvSpuSearchAPIResponse() *AlibabaIdleIsvSpuSearchAPIResponse {
	return poolAlibabaIdleIsvSpuSearchAPIResponse.Get().(*AlibabaIdleIsvSpuSearchAPIResponse)
}

// ReleaseAlibabaIdleIsvSpuSearchAPIResponse 将 AlibabaIdleIsvSpuSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvSpuSearchAPIResponse(v *AlibabaIdleIsvSpuSearchAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvSpuSearchAPIResponse.Put(v)
}
