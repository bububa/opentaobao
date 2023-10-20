package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptraveltradememoupdateAPIResponse 修改一笔交易备注 API返回值
// taobao.alitrip.travel.trade.memo.update
//
// 更新一笔交易备注
type TaobaoalitriptraveltradememoupdateAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptraveltradememoupdateAPIResponseModel
}

// TaobaoalitriptraveltradememoupdateAPIResponseModel is 修改一笔交易备注 成功返回结果
type TaobaoalitriptraveltradememoupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_memo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改交易备注返回值
	MemoUpdate *MemoUpdate `json:"memo_update,omitempty" xml:"memo_update,omitempty"`
}
