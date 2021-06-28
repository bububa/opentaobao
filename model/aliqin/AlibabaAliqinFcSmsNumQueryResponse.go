package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短信发送记录查询 APIResponse
alibaba.aliqin.fc.sms.num.query

短信发送记录查询。
*/
type AlibabaAliqinFcSmsNumQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcSmsNumQueryResponse `json:"alibaba_aliqin_fc_sms_num_query_response,omitempty"` 
    AlibabaAliqinFcSmsNumQueryResponse
}

/* model for simplify = false
type AlibabaAliqinFcSmsNumQueryResponse struct {

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // 每页数量
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 总量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 1
    
    Values  struct {
        FcPartnerSmsDetailDto  []FcPartnerSmsDetailDto `json:"fc_partner_sms_detail_dto,omitempty"`
    } `json:"values,omitempty"`
    

}
*/

type AlibabaAliqinFcSmsNumQueryResponse struct {

    // 当前页码
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 每页数量
    PageSize   int64 `json:"page_size,omitempty"`

    // 总量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 1
    Values   []FcPartnerSmsDetailDto `json:"values,omitempty"`

}
