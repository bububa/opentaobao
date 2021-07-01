package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkHrworkbenchCdpempsQueryAPIResponse
homs员工信息核对查询服务 API返回值
alibaba.wdk.hrworkbench.cdpemps.query

给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查 */
type AlibabaWdkHrworkbenchCdpempsQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkHrworkbenchCdpempsQueryAPIResponseModel
}

// AlibabaWdkHrworkbenchCdpempsQueryAPIResponseModel is homs员工信息核对查询服务 成功返回结果
type AlibabaWdkHrworkbenchCdpempsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_hrworkbench_cdpemps_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 数据集合
	Datas []AlibabaWdkHrworkbenchCdpempsQueryData `json:"datas,omitempty" xml:"datas>alibaba_wdk_hrworkbench_cdpemps_query_data,omitempty"`
	// 每一页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 业务code
	BizCode int64 `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 业务是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
