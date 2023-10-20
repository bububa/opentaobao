package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityGoodsQueryAPIResponse 营销抽奖-用户奖品查询 API返回值
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
type AlitripMerchantGalaxyActivityGoodsQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityGoodsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel is 营销抽奖-用户奖品查询 成功返回结果
type AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_goods_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityGoodsQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyActivityGoodsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityGoodsQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyActivityGoodsQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyActivityGoodsQueryAPIResponse
func GetAlitripMerchantGalaxyActivityGoodsQueryAPIResponse() *AlitripMerchantGalaxyActivityGoodsQueryAPIResponse {
	return poolAlitripMerchantGalaxyActivityGoodsQueryAPIResponse.Get().(*AlitripMerchantGalaxyActivityGoodsQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyActivityGoodsQueryAPIResponse 将 AlitripMerchantGalaxyActivityGoodsQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityGoodsQueryAPIResponse(v *AlitripMerchantGalaxyActivityGoodsQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityGoodsQueryAPIResponse.Put(v)
}
