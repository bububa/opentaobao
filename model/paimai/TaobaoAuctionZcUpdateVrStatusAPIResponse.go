package paimai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionZcUpdateVrStatusAPIResponse 如视VR更新活跃状态 API返回值
// taobao.auction.zc.update.vr.status
//
// 如视VR更新活跃状态
type TaobaoAuctionZcUpdateVrStatusAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionZcUpdateVrStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAuctionZcUpdateVrStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionZcUpdateVrStatusAPIResponseModel).Reset()
}

// TaobaoAuctionZcUpdateVrStatusAPIResponseModel is 如视VR更新活跃状态 成功返回结果
type TaobaoAuctionZcUpdateVrStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_zc_update_vr_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionZcUpdateVrStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAuctionZcUpdateVrStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionZcUpdateVrStatusAPIResponse)
	},
}

// GetTaobaoAuctionZcUpdateVrStatusAPIResponse 从 sync.Pool 获取 TaobaoAuctionZcUpdateVrStatusAPIResponse
func GetTaobaoAuctionZcUpdateVrStatusAPIResponse() *TaobaoAuctionZcUpdateVrStatusAPIResponse {
	return poolTaobaoAuctionZcUpdateVrStatusAPIResponse.Get().(*TaobaoAuctionZcUpdateVrStatusAPIResponse)
}

// ReleaseTaobaoAuctionZcUpdateVrStatusAPIResponse 将 TaobaoAuctionZcUpdateVrStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionZcUpdateVrStatusAPIResponse(v *TaobaoAuctionZcUpdateVrStatusAPIResponse) {
	v.Reset()
	poolTaobaoAuctionZcUpdateVrStatusAPIResponse.Put(v)
}
