package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionzcupdatevrstatusAPIResponse 如视VR更新活跃状态 API返回值
// taobao.auction.zc.update.vr.status
//
// 如视VR更新活跃状态
type TaobaoauctionzcupdatevrstatusAPIResponse struct {
	model.CommonResponse
	TaobaoauctionzcupdatevrstatusAPIResponseModel
}

// TaobaoauctionzcupdatevrstatusAPIResponseModel is 如视VR更新活跃状态 成功返回结果
type TaobaoauctionzcupdatevrstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_zc_update_vr_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4top `json:"result,omitempty" xml:"result,omitempty"`
}
