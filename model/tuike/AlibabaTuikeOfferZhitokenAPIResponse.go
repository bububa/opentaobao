package tuike

import (
	"encoding/xml"

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

// AlibabaTuikeOfferZhitokenAPIResponseModel is 生成阿里口令 成功返回结果
type AlibabaTuikeOfferZhitokenAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tuike_offer_zhitoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaTuikeOfferZhitokenResult `json:"result,omitempty" xml:"result,omitempty"`
}
