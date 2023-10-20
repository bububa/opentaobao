package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkShopRecommendGetAPIResponse 淘宝客-公用-店铺关联推荐 API返回值
// taobao.tbk.shop.recommend.get
//
// 入参卖家id，可推荐与此店铺相关联的相关店铺。
type TaobaoTbkShopRecommendGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkShopRecommendGetAPIResponseModel
}

// TaobaoTbkShopRecommendGetAPIResponseModel is 淘宝客-公用-店铺关联推荐 成功返回结果
type TaobaoTbkShopRecommendGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_shop_recommend_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝客店铺
	Results []NtbkShop `json:"results,omitempty" xml:"results>ntbk_shop,omitempty"`
}
