package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScPunishOrderGetAPIResponse 淘宝客-服务商-处罚订单查询 API返回值
// taobao.tbk.sc.punish.order.get
//
// 服务商使用。可通过该接口查询推广者账号下在处罚管理后台，可直接下载的处罚订单明细。请注意：若是基于账号(member)、媒体id(site)、推广位(adzone)、渠道方(relationid)维度的完整处罚，对应订单明细请根据处罚后台对应的处罚订单时间说明，在“推广订单明细”中筛选对应订单。
type TaobaoTbkScPunishOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScPunishOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScPunishOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScPunishOrderGetAPIResponseModel).Reset()
}

// TaobaoTbkScPunishOrderGetAPIResponseModel is 淘宝客-服务商-处罚订单查询 成功返回结果
type TaobaoTbkScPunishOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_punish_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询的对象
	Result *TaobaoTbkScPunishOrderGetRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScPunishOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTbkScPunishOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPunishOrderGetAPIResponse)
	},
}

// GetTaobaoTbkScPunishOrderGetAPIResponse 从 sync.Pool 获取 TaobaoTbkScPunishOrderGetAPIResponse
func GetTaobaoTbkScPunishOrderGetAPIResponse() *TaobaoTbkScPunishOrderGetAPIResponse {
	return poolTaobaoTbkScPunishOrderGetAPIResponse.Get().(*TaobaoTbkScPunishOrderGetAPIResponse)
}

// ReleaseTaobaoTbkScPunishOrderGetAPIResponse 将 TaobaoTbkScPunishOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScPunishOrderGetAPIResponse(v *TaobaoTbkScPunishOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkScPunishOrderGetAPIResponse.Put(v)
}
