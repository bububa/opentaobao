package tuike

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTuikeOfferGetAPIResponse 推广商品查询接口 API返回值
// alibaba.tuike.offer.get
//
// 查询1688推客平台卖家推广中的商品信息
type AlibabaTuikeOfferGetAPIResponse struct {
	model.CommonResponse
	AlibabaTuikeOfferGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTuikeOfferGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTuikeOfferGetAPIResponseModel).Reset()
}

// AlibabaTuikeOfferGetAPIResponseModel is 推广商品查询接口 成功返回结果
type AlibabaTuikeOfferGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tuike_offer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果模型
	Result *TaOfferSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTuikeOfferGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTuikeOfferGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTuikeOfferGetAPIResponse)
	},
}

// GetAlibabaTuikeOfferGetAPIResponse 从 sync.Pool 获取 AlibabaTuikeOfferGetAPIResponse
func GetAlibabaTuikeOfferGetAPIResponse() *AlibabaTuikeOfferGetAPIResponse {
	return poolAlibabaTuikeOfferGetAPIResponse.Get().(*AlibabaTuikeOfferGetAPIResponse)
}

// ReleaseAlibabaTuikeOfferGetAPIResponse 将 AlibabaTuikeOfferGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTuikeOfferGetAPIResponse(v *AlibabaTuikeOfferGetAPIResponse) {
	v.Reset()
	poolAlibabaTuikeOfferGetAPIResponse.Put(v)
}
