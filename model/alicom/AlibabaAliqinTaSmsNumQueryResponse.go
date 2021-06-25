package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短信查询 APIResponse
alibaba.aliqin.ta.sms.num.query

查询短信发送揭露
*/
type AlibabaAliqinTaSmsNumQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinTaSmsNumQueryResponse `json:"alibaba_aliqin_ta_sms_num_query_response,omitempty"`
}

type AlibabaAliqinTaSmsNumQueryResponse struct {

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
