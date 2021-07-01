package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyStoreQueryAPIResponse
根据门店查分类信息 API返回值
taobao.omniitem.classify.store.query

根据门店查分类信息 */
type TaobaoOmniitemClassifyStoreQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemClassifyStoreQueryAPIResponseModel
}

// TaobaoOmniitemClassifyStoreQueryAPIResponseModel is 根据门店查分类信息 成功返回结果
type TaobaoOmniitemClassifyStoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_classify_store_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemClassifyStoreQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
