package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripGrouptourProductUploadAPIResponse 新版跟团游商品维护接口 API返回值
// alitrip.grouptour.product.upload
//
// 新版跟团游商品维护接口
type AlitripGrouptourProductUploadAPIResponse struct {
	model.CommonResponse
	AlitripGrouptourProductUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripGrouptourProductUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripGrouptourProductUploadAPIResponseModel).Reset()
}

// AlitripGrouptourProductUploadAPIResponseModel is 新版跟团游商品维护接口 成功返回结果
type AlitripGrouptourProductUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_grouptour_product_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// firstResult
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripGrouptourProductUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripGrouptourProductUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripGrouptourProductUploadAPIResponse)
	},
}

// GetAlitripGrouptourProductUploadAPIResponse 从 sync.Pool 获取 AlitripGrouptourProductUploadAPIResponse
func GetAlitripGrouptourProductUploadAPIResponse() *AlitripGrouptourProductUploadAPIResponse {
	return poolAlitripGrouptourProductUploadAPIResponse.Get().(*AlitripGrouptourProductUploadAPIResponse)
}

// ReleaseAlitripGrouptourProductUploadAPIResponse 将 AlitripGrouptourProductUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripGrouptourProductUploadAPIResponse(v *AlitripGrouptourProductUploadAPIResponse) {
	v.Reset()
	poolAlitripGrouptourProductUploadAPIResponse.Put(v)
}
