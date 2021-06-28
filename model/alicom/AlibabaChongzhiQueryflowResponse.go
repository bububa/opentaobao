package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的流量宝贝接口 APIResponse
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝
*/
type AlibabaChongzhiQueryflowAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_chongzhi_queryflow_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   int64 `json:"result,omitempty" xml:"