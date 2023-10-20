package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradetoolsitemsunbindAPIResponse 交易开放商品解绑 API返回值
// taobao.opentrade.tools.items.unbind
//
// 交易开放商品解绑
type TaobaoopentradetoolsitemsunbindAPIResponse struct {
	model.CommonResponse
	TaobaoopentradetoolsitemsunbindAPIResponseModel
}

// TaobaoopentradetoolsitemsunbindAPIResponseModel is 交易开放商品解绑 成功返回结果
type TaobaoopentradetoolsitemsunbindAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_tools_items_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 解绑返回结构
	Results []ItemUnBindResult `json:"results,omitempty" xml:"results>item_un_bind_result,omitempty"`
}
