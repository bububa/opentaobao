package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemsMarkingAPIResponse
商品通自动打标 API返回值
taobao.qimen.items.marking

调用该接口，对商品进行XXXX标的打标、去标的动作。 */
type TaobaoQimenItemsMarkingAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemsMarkingAPIResponseModel
}

// TaobaoQimenItemsMarkingAPIResponseModel is 商品通自动打标 成功返回结果
type TaobaoQimenItemsMarkingAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_items_marking_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// flag
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
