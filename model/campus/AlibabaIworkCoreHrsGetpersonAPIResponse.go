package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIworkCoreHrsGetpersonAPIResponse
获取神鲸用户基本信息 API返回值
alibaba.iwork.core.hrs.getperson

神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询 */
type AlibabaIworkCoreHrsGetpersonAPIResponse struct {
	model.CommonResponse
	AlibabaIworkCoreHrsGetpersonAPIResponseModel
}

// AlibabaIworkCoreHrsGetpersonAPIResponseModel is 获取神鲸用户基本信息 成功返回结果
type AlibabaIworkCoreHrsGetpersonAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_iwork_core_hrs_getperson_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
