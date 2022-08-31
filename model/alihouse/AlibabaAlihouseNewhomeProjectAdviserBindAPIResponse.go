package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse 置业顾问批量绑定楼盘 API返回值
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
type AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel is 置业顾问批量绑定楼盘 成功返回结果
type AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectAdviserBindResult `json:"result,omitempty" xml:"result,omitempty"`
}
