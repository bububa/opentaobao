package paimai

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAuctionZcMerchantUserCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionZcMerchantUserCheckAPIResponseModel).Reset()
}

// TaobaoAuctionZcMerchantUserCheckAPIResponseModel is 通过手机号确认阿里资产商家 成功返回结果
type TaobaoAuctionZcMerchantUserCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_zc_merchant_user_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionZcMerchantUserCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAuctionZcMerchantUserCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionZcMerchantUserCheckAPIResponse)
	},
}

// GetTaobaoAuctionZcMerchantUserCheckAPIResponse 从 sync.Pool 获取 TaobaoAuctionZcMerchantUserCheckAPIResponse
func GetTaobaoAuctionZcMerchantUserCheckAPIResponse() *TaobaoAuctionZcMerchantUserCheckAPIResponse {
	return poolTaobaoAuctionZcMerchantUserCheckAPIResponse.Get().(*TaobaoAuctionZcMerchantUserCheckAPIResponse)
}

// ReleaseTaobaoAuctionZcMerchantUserCheckAPIResponse 将 TaobaoAuctionZcMerchantUserCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionZcMerchantUserCheckAPIResponse(v *TaobaoAuctionZcMerchantUserCheckAPIResponse) {
	v.Reset()
	poolTaobaoAuctionZcMerchantUserCheckAPIResponse.Put(v)
}
