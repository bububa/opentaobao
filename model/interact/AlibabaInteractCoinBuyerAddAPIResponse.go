package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractCoinBuyerAddAPIResponse
平台向买家发放淘金币 API返回值
alibaba.interact.coin.buyer.add

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。ISV调用该接口向买家发放平台淘金币，需要优惠平台运营审核ISV资质。 */
type AlibabaInteractCoinBuyerAddAPIResponse struct {
	model.CommonResponse
	AlibabaInteractCoinBuyerAddAPIResponseModel
}

// AlibabaInteractCoinBuyerAddAPIResponseModel is 平台向买家发放淘金币 成功返回结果
type AlibabaInteractCoinBuyerAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_coin_buyer_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
