package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改类目 APIResponse
alibaba.wdk.item.category.update

修改类目
*/
type AlibabaWdkItemCategoryUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_item_category_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkItemCategoryUpdateResult `json:"result,omitempty" xml:"