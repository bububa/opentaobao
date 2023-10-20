package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradetoolsitemsbindAPIResponse 交易开放商品绑定 API返回值
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
type TaobaoopentradetoolsitemsbindAPIResponse struct {
	model.CommonResponse
	TaobaoopentradetoolsitemsbindAPIResponseModel
}

// TaobaoopentradetoolsitemsbindAPIResponseModel is 交易开放商品绑定 成功返回结果
type TaobaoopentradetoolsitemsbindAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_tools_items_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定返回结构
	Results []ItemBindResult `json:"results,omitempty" xml:"results>item_bind_result,omitempty"`
}
