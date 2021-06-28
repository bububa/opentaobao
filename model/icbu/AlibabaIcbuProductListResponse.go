package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询 APIResponse
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。
*/
type AlibabaIcbuProductListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总数
    
    TotalItem   int64 `json:"total_item,omitempty" xml:"