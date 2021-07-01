package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签绑定接口 API返回值 
taobao.uscesl.biz.esl.bind

电子价签商品绑定接口
*/
type TaobaoUsceslBizEslBindAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizEslBindAPIResponseModel
}

// 电子价签绑定接口 成功返回结果
type TaobaoUsceslBizEslBindAPIResponseModel struct {
    XMLName xml.Name `xml:"uscesl_biz_esl_bind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 成功与否看result.success，返回true或者false
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
