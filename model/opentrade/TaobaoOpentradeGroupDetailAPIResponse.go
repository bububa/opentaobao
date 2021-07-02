package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupDetailAPIResponse 组团购场景查询团详情 API返回值
// taobao.opentrade.group.detail
//
// 组团购场景下，查询团详情
type TaobaoOpentradeGroupDetailAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeGroupDetailAPIResponseModel
}

// TaobaoOpentradeGroupDetailAPIResponseModel is 组团购场景查询团详情 成功返回结果
type TaobaoOpentradeGroupDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_group_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 团信息
	Result *GroupDetailResponse `json:"result,omitempty" xml:"result,omitempty"`
}
