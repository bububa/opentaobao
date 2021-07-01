package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupJoinAPIResponse
组团购场景参团 API返回值
taobao.opentrade.group.join

组团购场景下，用户参团 */
type TaobaoOpentradeGroupJoinAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeGroupJoinAPIResponseModel
}

// TaobaoOpentradeGroupJoinAPIResponseModel is 组团购场景参团 成功返回结果
type TaobaoOpentradeGroupJoinAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_group_join_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
