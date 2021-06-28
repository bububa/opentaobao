package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站图片银行查询接口 APIResponse
alibaba.icbu.photobank.list

国际站图片银行查询接口
*/
type AlibabaIcbuPhotobankListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_photobank_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 图片查询结果
    
    PaginationQueryList   *PaginationQueryList `json:"pagination_query_list,omitempty" xml:"