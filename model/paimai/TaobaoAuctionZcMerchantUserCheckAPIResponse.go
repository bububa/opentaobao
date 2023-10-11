package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionZcMerchantUserCheckAPIResponse 通过手机号确认阿里资产商家 API返回值
// taobao.auction.zc.merchant.user.check
//
// 通过手机号确认阿里资产商家
type TaobaoAuctionZcMerchantUserCheckAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionZcMerchantUserCheckAPIResponseModel
}

// TaobaoAuctionZcMerchantUserCheckAPIResponseModel is 通过手机号确认阿里资产商家 成功返回结果
type TaobaoAuctionZcMerchantUserCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_zc_merchant_user_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}
