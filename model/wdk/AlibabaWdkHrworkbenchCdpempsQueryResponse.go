package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
homs员工信息核对查询服务 APIResponse
alibaba.wdk.hrworkbench.cdpemps.query

给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查
*/
type AlibabaWdkHrworkbenchCdpempsQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_hrworkbench_cdpemps_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回信息
    
    Message   string `json:"message,omitempty" xml:"