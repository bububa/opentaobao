package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
homs人力资源组织树查询接口 APIResponse
alibaba.wdk.hrworkbench.cdporgs.query

提供查询homs人力组织树的接口，按照商家做权限隔离。
*/
type AlibabaWdkHrworkbenchCdporgsQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkHrworkbenchCdporgsQueryResponse `json:"alibaba_wdk_hrworkbench_cdporgs_query_response,omitempty"`
}

type AlibabaWdkHrworkbenchCdporgsQueryResponse struct {

    // 结果信息
    Message   string `json:"message,omitempty"`

    // 鹰眼id
    TraceId   string `json:"trace_id,omitempty"`

    // 数据集合
    Datas   []AlibabaWdkHrworkbenchCdporgsQueryData `json:"datas,omitempty"`

    // 参数code
    BizCode   int64 `json:"biz_code,omitempty"`

    // 请求是否成功
    BizSuccess   bool `json:"biz_success,omitempty"`

}
