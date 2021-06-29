package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目信息 APIResponse
alibaba.wholesale.category.get

获取类目信息
*/
type AlibabaWholesaleCategoryGetAPIResponse struct {
    model.CommonResponse
    AlibabaWholesaleCategoryGetResponse
}

type AlibabaWholesaleCategoryGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wholesale_category_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 类目结果
    
    WholesaleCategoryResult   *WholesaleCategoryOpenResult `json:"wholesale_category_result,omitempty" xml:"wholesale_category_result,omitempty"`

    
}
