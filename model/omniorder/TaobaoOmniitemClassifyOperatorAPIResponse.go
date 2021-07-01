package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加/修改分类 API返回值 
taobao.omniitem.classify.operator

添加/修改分类
*/
type TaobaoOmniitemClassifyOperatorAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyOperatorAPIResponseModel
}

// 添加/修改分类 成功返回结果
type TaobaoOmniitemClassifyOperatorAPIResponseModel struct {
    XMLName xml.Name `xml:"omniitem_classify_operator_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniitemClassifyOperatorResult `json:"result,omitempty" xml:"result,omitempty"`
}
