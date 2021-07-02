package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemClassifyQueryAPIResponse 查询分类信息 API返回值
// taobao.omniitem.classify.query
//
// 通过查询关键字，分页查询分类信息
type TaobaoOmniitemClassifyQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemClassifyQueryAPIResponseModel
}

// TaobaoOmniitemClassifyQueryAPIResponseModel is 查询分类信息 成功返回结果
type TaobaoOmniitemClassifyQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_classify_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemClassifyQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
