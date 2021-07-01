package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse
提交预售证 API返回值
alibaba.alihouse.newhome.project.presalepermit.submit

提交楼盘预售证信息 */
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel is 提交预售证 成功返回结果
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_presalepermit_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
