package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU类目树获取接口 APIResponse
alibaba.icbu.category.get.new

获取商品发布类目
*/
type AlibabaIcbuCategoryGetNewAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_category_get_new_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 类目信息
    
    Category   *PostCategory `json:"category,omitempty" xml:"