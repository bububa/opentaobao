package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIResponse
码核查状态同步-医保 API返回值
alibaba.alihealth.drug.code.code.check.medical.insurance

服务描述
核查平台在进行医保单据鉴证核查时，会记录单据中所有提交的追溯码信息；单据中的
追溯码包含所有正常和异常的数据；
此接口，针对正式鉴核的单据中提交的有效的、正常状态的追溯码，提供可由核查平台
发起，按单据鉴核时间顺序组织，向码上放心平台同步每笔单据核销的码状态信息；
入参采用数组方式提供，一次同步最多支持100条记录 */
type AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIResponseModel
}

// AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIResponseModel is 码核查状态同步-医保 成功返回结果
type AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_code_check_medical_insurance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
