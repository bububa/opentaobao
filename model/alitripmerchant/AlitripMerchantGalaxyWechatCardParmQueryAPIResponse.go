package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatCardParmQueryAPIResponse 微信会员卡添加 API返回值
// alitrip.merchant.galaxy.wechat.card.parm.query
//
// 微信会员卡添加参数获取
type AlitripMerchantGalaxyWechatCardParmQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatCardParmQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel is 微信会员卡添加 成功返回结果
type AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_card_parm_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyWechatCardParmQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatCardParmQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatCardParmQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatCardParmQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatCardParmQueryAPIResponse
func GetAlitripMerchantGalaxyWechatCardParmQueryAPIResponse() *AlitripMerchantGalaxyWechatCardParmQueryAPIResponse {
	return poolAlitripMerchantGalaxyWechatCardParmQueryAPIResponse.Get().(*AlitripMerchantGalaxyWechatCardParmQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatCardParmQueryAPIResponse 将 AlitripMerchantGalaxyWechatCardParmQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatCardParmQueryAPIResponse(v *AlitripMerchantGalaxyWechatCardParmQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatCardParmQueryAPIResponse.Put(v)
}
