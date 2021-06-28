package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
homs员工信息核对查询服务 APIResponse
alibaba.wdk.hrworkbench.cdpemps.query

给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查
*/
type AlibabaWdkHrworkbenchCdpempsQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkHrworkbenchCdpempsQueryResponse `json:"alibaba_wdk_hrworkbench_cdpemps_query_response,omitempty"` 
    AlibabaWdkHrworkbenchCdpempsQueryResponse
}

/* model for simplify = false
type AlibabaWdkHrworkbenchCdpempsQueryResponse struct {

    // 返回信息
    
    Message   string `json:"message,omitempty"`
    

    // 总页数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 鹰眼id
    
    TraceId   string `json:"trace_id,omitempty"`
    

    // 数据集合
    
    Datas  struct {
        AlibabaWdkHrworkbenchCdpempsQueryData  []AlibabaWdkHrworkbenchCdpempsQueryData `json:"alibaba_wdk_hrworkbench_cdpemps_query_data,omitempty"`
    } `json:"datas,omitempty"`
    

    // 每一页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 业务code
    
    BizCode   int64 `json:"biz_code,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 业务是否成功
    
    BizSuccess   bool `json:"biz_success,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

}
*/

type AlibabaWdkHrworkbenchCdpempsQueryResponse struct {

    // 返回信息
    Message   string `json:"message,omitempty"`

    // 总页数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 鹰眼id
    TraceId   string `json:"trace_id,omitempty"`

    // 数据集合
    Datas   []AlibabaWdkHrworkbenchCdpempsQueryData `json:"datas,omitempty"`

    // 每一页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 业务code
    BizCode   int64 `json:"biz_code,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 业务是否成功
    BizSuccess   bool `json:"biz_success,omitempty"`

    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty"`

}
