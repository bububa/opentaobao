package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
homs人力资源组织树查询接口 APIResponse
alibaba.wdk.hrworkbench.cdporgs.query

提供查询homs人力组织树的接口，按照商家做权限隔离。
*/
type AlibabaWdkHrworkbenchCdporgsQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_hrworkbench_cdporgs_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Message   string `json:"message,omitempty" xml:"