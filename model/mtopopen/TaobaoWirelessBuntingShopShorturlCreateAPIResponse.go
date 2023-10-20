package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowirelessbuntingshopshorturlcreateAPIResponse 通过店铺id取得短链 API返回值
// taobao.wireless.bunting.shop.shorturl.create
//
// 通过店铺id取得短链
type TaobaowirelessbuntingshopshorturlcreateAPIResponse struct {
	model.CommonResponse
	TaobaowirelessbuntingshopshorturlcreateAPIResponseModel
}

// TaobaowirelessbuntingshopshorturlcreateAPIResponseModel is 通过店铺id取得短链 成功返回结果
type TaobaowirelessbuntingshopshorturlcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"wireless_bunting_shop_shorturl_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短链
	Shorturl string `json:"shorturl,omitempty" xml:"shorturl,omitempty"`
}
