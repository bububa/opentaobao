package dengta

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaWxaccountPriceChangeAPIResponse
微信公众号价格变化通知 API返回值
alibaba.pictures.dengta.wxaccount.price.change

微信公众号推广价格变更通知接口 */
type AlibabaPicturesDengtaWxaccountPriceChangeAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaWxaccountPriceChangeAPIResponseModel
}

// AlibabaPicturesDengtaWxaccountPriceChangeAPIResponseModel is 微信公众号价格变化通知 成功返回结果
type AlibabaPicturesDengtaWxaccountPriceChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_wxaccount_price_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}
