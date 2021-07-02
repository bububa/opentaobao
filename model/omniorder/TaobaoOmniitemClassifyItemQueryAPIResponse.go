package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemClassifyItemQueryAPIResponse 根据分类查商品信息 API返回值
// taobao.omniitem.classify.item.query
//
// 商家根据分类查商品
type TaobaoOmniitemClassifyItemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemClassifyItemQueryAPIResponseModel
}

// TaobaoOmniitemClassifyItemQueryAPIResponseModel is 根据分类查商品信息 成功返回结果
type TaobaoOmniitemClassifyItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_classify_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemClassifyItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
