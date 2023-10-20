package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenShopSynchronizeAPIResponse 店铺同步接口 API返回值
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
type TaobaoQimenShopSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenShopSynchronizeAPIResponseModel
}

// TaobaoQimenShopSynchronizeAPIResponseModel is 店铺同步接口 成功返回结果
type TaobaoQimenShopSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_shop_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Response
	Response *TaobaoQimenShopSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
