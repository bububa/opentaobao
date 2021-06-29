package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布类目获取 API返回值 
alibaba.icbu.category.get

获取商品发布类目
*/
type AlibabaIcbuCategoryGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuCategoryGetResponse
}

// 商品发布类目获取 成功返回结果
type AlibabaIcbuCategoryGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_category_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 类目信息
    Category   *PostCategory `json:"category,omitempty" xml:"category,omitempty"`
}
