package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemItemDeleteAPIResponse
全渠道商品删除 API返回值
taobao.omniitem.item.delete

全渠道商品删除，能够对门店商品库商品进行删除动作 */
type TaobaoOmniitemItemDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemItemDeleteAPIResponseModel
}

// TaobaoOmniitemItemDeleteAPIResponseModel is 全渠道商品删除 成功返回结果
type TaobaoOmniitemItemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemItemDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
