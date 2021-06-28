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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wholesale_category_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 类目结果
    
    WholesaleCategoryResult   *WholesaleCategoryOpenResult `json:"wholesale_category_result,omitempty" xml:"