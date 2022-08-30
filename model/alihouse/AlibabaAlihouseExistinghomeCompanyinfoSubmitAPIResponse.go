package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse 进件接口 API返回值
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
type AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel
}

// AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel is 进件接口 成功返回结果
type AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_companyinfo_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeCompanyinfoSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
