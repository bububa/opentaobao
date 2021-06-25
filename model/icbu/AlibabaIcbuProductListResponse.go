package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品查询 APIResponse
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。
*/
type AlibabaIcbuProductListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuProductListResponse `json:"alibaba_icbu_product_list_response,omitempty"`
}

type AlibabaIcbuProductListResponse struct {

    // 总数
    TotalItem   int64 `json:"total_item,omitempty"`

    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 每页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 商品概要信息列表
    Products   []AlibabaProductBriefResponse `json:"products,omitempty"`

}
