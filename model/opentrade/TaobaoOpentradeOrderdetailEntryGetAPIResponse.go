package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeOrderdetailEntryGetAPIResponse 获取订单详情页跳转入口配置 API返回值
// taobao.opentrade.orderdetail.entry.get
//
// 获取订单详情页跳转入口配置
type TaobaoOpentradeOrderdetailEntryGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeOrderdetailEntryGetAPIResponseModel
}

// TaobaoOpentradeOrderdetailEntryGetAPIResponseModel is 获取订单详情页跳转入口配置 成功返回结果
type TaobaoOpentradeOrderdetailEntryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_orderdetail_entry_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 跳转入口配置
	Result []OrderDetailEntryConfig `json:"result,omitempty" xml:"result>order_detail_entry_config,omitempty"`
}
