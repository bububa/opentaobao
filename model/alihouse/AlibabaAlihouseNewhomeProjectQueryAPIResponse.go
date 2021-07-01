package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectQueryAPIResponse
查询楼盘相关信息 API返回值
alibaba.alihouse.newhome.project.query

根据outerid查询楼盘相关信息 */
type AlibabaAlihouseNewhomeProjectQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectQueryAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectQueryAPIResponseModel is 查询楼盘相关信息 成功返回结果
type AlibabaAlihouseNewhomeProjectQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
