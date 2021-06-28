package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
国际站图片银行查询接口 APIResponse
alibaba.icbu.photobank.list

国际站图片银行查询接口
*/
type AlibabaIcbuPhotobankListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuPhotobankListResponse `json:"alibaba_icbu_photobank_list_response,omitempty"` 
    AlibabaIcbuPhotobankListResponse
}

/* model for simplify = false
type AlibabaIcbuPhotobankListResponse struct {

    // 图片查询结果
    
    PaginationQueryList  *struct {
        PaginationQueryList  *PaginationQueryList `json:"pagination_query_list,omitempty"`
    } `json:"pagination_query_list,omitempty"`
    

}
*/

type AlibabaIcbuPhotobankListResponse struct {

    // 图片查询结果
    PaginationQueryList   *PaginationQueryList `json:"pagination_query_list,omitempty"`

}
