package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetdruglicenseAPIResponse
获取药品资质信息 API返回值
alibaba.alihealth.drug.kyt.getdruglicense

获取药品的资质信息。 */
type AlibabaAlihealthDrugKytGetdruglicenseAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetdruglicenseAPIResponseModel
}

// AlibabaAlihealthDrugKytGetdruglicenseAPIResponseModel is 获取药品资质信息 成功返回结果
type AlibabaAlihealthDrugKytGetdruglicenseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getdruglicense_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytGetdruglicenseResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
