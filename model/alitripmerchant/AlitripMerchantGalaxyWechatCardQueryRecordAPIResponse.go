package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse 微信会员卡领取记录查询 API返回值
// alitrip.merchant.galaxy.wechat.card.query.record
//
// 微信会员卡领取记录查询
type AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatCardQueryRecordAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatCardQueryRecordAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatCardQueryRecordAPIResponseModel is 微信会员卡领取记录查询 成功返回结果
type AlitripMerchantGalaxyWechatCardQueryRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_card_query_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyWechatCardQueryRecordResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatCardQueryRecordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse
func GetAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse() *AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse {
	return poolAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse.Get().(*AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse 将 AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse(v *AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatCardQueryRecordAPIResponse.Put(v)
}
