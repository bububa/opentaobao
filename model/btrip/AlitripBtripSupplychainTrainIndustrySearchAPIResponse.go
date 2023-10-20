package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainIndustrySearchAPIResponse 火车票行业搜索接口 API返回值
// alitrip.btrip.supplychain.train.industry.search
//
// 【商旅】火车票行业搜索接口
type AlitripBtripSupplychainTrainIndustrySearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainTrainIndustrySearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainIndustrySearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainTrainIndustrySearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainTrainIndustrySearchAPIResponseModel is 火车票行业搜索接口 成功返回结果
type AlitripBtripSupplychainTrainIndustrySearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_industry_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainIndustrySearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainTrainIndustrySearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainTrainIndustrySearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainTrainIndustrySearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainTrainIndustrySearchAPIResponse
func GetAlitripBtripSupplychainTrainIndustrySearchAPIResponse() *AlitripBtripSupplychainTrainIndustrySearchAPIResponse {
	return poolAlitripBtripSupplychainTrainIndustrySearchAPIResponse.Get().(*AlitripBtripSupplychainTrainIndustrySearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainTrainIndustrySearchAPIResponse 将 AlitripBtripSupplychainTrainIndustrySearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainTrainIndustrySearchAPIResponse(v *AlitripBtripSupplychainTrainIndustrySearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainTrainIndustrySearchAPIResponse.Put(v)
}
