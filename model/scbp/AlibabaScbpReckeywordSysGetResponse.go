package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系统推荐 APIResponse
alibaba.scbp.reckeyword.sys.get

查询系统推荐词
*/
type AlibabaScbpReckeywordSysGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_reckeyword_sys_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"