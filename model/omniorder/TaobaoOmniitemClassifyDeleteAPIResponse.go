package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemClassifyDeleteAPIResponse 删除一个分类 API返回值
// taobao.omniitem.classify.delete
//
// 删除一个分类
type TaobaoOmniitemClassifyDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemClassifyDeleteAPIResponseModel
}

// TaobaoOmniitemClassifyDeleteAPIResponseModel is 删除一个分类 成功返回结果
type TaobaoOmniitemClassifyDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_classify_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemClassifyDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
