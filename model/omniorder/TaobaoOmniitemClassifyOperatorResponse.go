package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加/修改分类 APIResponse
taobao.omniitem.classify.operator

添加/修改分类
*/
type TaobaoOmniitemClassifyOperatorAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyOperatorResponse
}

type TaobaoOmniitemClassifyOperatorResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_operator_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemClassifyOperatorResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
