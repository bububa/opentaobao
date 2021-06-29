package aliexpresssumaitong

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关务所需的申报清关字段 APIResponse
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段
*/
type AliexpressTaxationCalculateOpenQueryAPIResponse struct {
    model.CommonResponse
    AliexpressTaxationCalculateOpenQueryResponse
}

type AliexpressTaxationCalculateOpenQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_taxation_calculate_open_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ResponseDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
