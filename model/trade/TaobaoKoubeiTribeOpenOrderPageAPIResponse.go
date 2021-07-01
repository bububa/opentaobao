package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiTribeOpenOrderPageAPIResponse
口碑综合体订单列表信息查询 API返回值
taobao.koubei.tribe.open.order.page

查询口碑商圈用户的订单列表信息 */
type TaobaoKoubeiTribeOpenOrderPageAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiTribeOpenOrderPageAPIResponseModel
}

// TaobaoKoubeiTribeOpenOrderPageAPIResponseModel is 口碑综合体订单列表信息查询 成功返回结果
type TaobaoKoubeiTribeOpenOrderPageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_tribe_open_order_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiTribeOpenOrderPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
