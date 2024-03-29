package admarket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAdmarketAdBidAPIResponse 广告竞价服务 API返回值
// yunos.admarket.ad.bid
//
// 广告竞价服务
type YunosAdmarketAdBidAPIResponse struct {
	model.CommonResponse
	YunosAdmarketAdBidAPIResponseModel
}

// Reset 清空结构体
func (m *YunosAdmarketAdBidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosAdmarketAdBidAPIResponseModel).Reset()
}

// YunosAdmarketAdBidAPIResponseModel is 广告竞价服务 成功返回结果
type YunosAdmarketAdBidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_admarket_ad_bid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 返回结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 响应码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回结果
	Result *BidResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 是否操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YunosAdmarketAdBidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Id = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.Result = nil
	m.IsSuccess = false
}

var poolYunosAdmarketAdBidAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosAdmarketAdBidAPIResponse)
	},
}

// GetYunosAdmarketAdBidAPIResponse 从 sync.Pool 获取 YunosAdmarketAdBidAPIResponse
func GetYunosAdmarketAdBidAPIResponse() *YunosAdmarketAdBidAPIResponse {
	return poolYunosAdmarketAdBidAPIResponse.Get().(*YunosAdmarketAdBidAPIResponse)
}

// ReleaseYunosAdmarketAdBidAPIResponse 将 YunosAdmarketAdBidAPIResponse 保存到 sync.Pool
func ReleaseYunosAdmarketAdBidAPIResponse(v *YunosAdmarketAdBidAPIResponse) {
	v.Reset()
	poolYunosAdmarketAdBidAPIResponse.Put(v)
}
