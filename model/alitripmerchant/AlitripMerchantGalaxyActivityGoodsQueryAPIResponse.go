package alitripmerchant

import (
	"encoding/xml"

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

// AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel is 营销抽奖-用户奖品查询 成功返回结果
type AlitripMerchantGalaxyActivityGoodsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_goods_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityGoodsQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
