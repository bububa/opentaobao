package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgPunishOrderGetAPIResponse
淘宝客-推广者-处罚订单查询 API返回值
taobao.tbk.dg.punish.order.get

新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的 */
type TaobaoTbkDgPunishOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgPunishOrderGetAPIResponseModel
}

// TaobaoTbkDgPunishOrderGetAPIResponseModel is 淘宝客-推广者-处罚订单查询 成功返回结果
type TaobaoTbkDgPunishOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_punish_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询的对象
	Result *TaobaoTbkDgPunishOrderGetRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}
