package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增电子价签商家 APIResponse
taobao.uscesl.biz.brand.insert

一个电子价签业务身份下新增商家接口
*/
type TaobaoUsceslBizBrandInsertAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizBrandInsertResponse
}

type TaobaoUsceslBizBrandInsertResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_brand_insert_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
