package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse 案场活动楼盘维护 API返回值
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel
}

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel is 案场活动楼盘维护 成功返回结果
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_casefield_activity_project_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
