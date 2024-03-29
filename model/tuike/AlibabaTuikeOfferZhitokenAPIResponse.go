package tuike

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTuikeOfferZhitokenAPIResponse 生成阿里口令 API返回值
// alibaba.tuike.offer.zhitoken
//
// 推荐链接生产吱口令
type AlibabaTuikeOfferZhitokenAPIResponse struct {
	model.CommonResponse
	AlibabaTuikeOfferZhitokenAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTuikeOfferZhitokenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTuikeOfferZhitokenAPIResponseModel).Reset()
}

// AlibabaTuikeOfferZhitokenAPIResponseModel is 生成阿里口令 成功返回结果
type AlibabaTuikeOfferZhitokenAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tuike_offer_zhitoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaTuikeOfferZhitokenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTuikeOfferZhitokenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTuikeOfferZhitokenAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTuikeOfferZhitokenAPIResponse)
	},
}

// GetAlibabaTuikeOfferZhitokenAPIResponse 从 sync.Pool 获取 AlibabaTuikeOfferZhitokenAPIResponse
func GetAlibabaTuikeOfferZhitokenAPIResponse() *AlibabaTuikeOfferZhitokenAPIResponse {
	return poolAlibabaTuikeOfferZhitokenAPIResponse.Get().(*AlibabaTuikeOfferZhitokenAPIResponse)
}

// ReleaseAlibabaTuikeOfferZhitokenAPIResponse 将 AlibabaTuikeOfferZhitokenAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTuikeOfferZhitokenAPIResponse(v *AlibabaTuikeOfferZhitokenAPIResponse) {
	v.Reset()
	poolAlibabaTuikeOfferZhitokenAPIResponse.Put(v)
}
