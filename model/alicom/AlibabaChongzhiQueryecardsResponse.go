package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的话费宝贝接口 APIResponse
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝
*/
type AlibabaChongzhiQueryecardsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_chongzhi_queryecards_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   int64 `json:"result,omitempty" xml:"