package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkHrworkbenchCdporgsQueryAPIResponse homs人力资源组织树查询接口 API返回值
// alibaba.wdk.hrworkbench.cdporgs.query
//
// 提供查询homs人力组织树的接口，按照商家做权限隔离。
type AlibabaWdkHrworkbenchCdporgsQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkHrworkbenchCdporgsQueryAPIResponseModel
}

// AlibabaWdkHrworkbenchCdporgsQueryAPIResponseModel is homs人力资源组织树查询接口 成功返回结果
type AlibabaWdkHrworkbenchCdporgsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_hrworkbench_cdporgs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据集合
	Datas []AlibabaWdkHrworkbenchCdporgsQueryData `json:"datas,omitempty" xml:"datas>alibaba_wdk_hrworkbench_cdporgs_query_data,omitempty"`
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 参数code
	BizCode int64 `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 请求是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
