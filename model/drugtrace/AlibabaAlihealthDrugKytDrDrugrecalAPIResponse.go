package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrDrugrecalAPIResponse
疫苗药品召回 API返回值
alibaba.alihealth.drug.kyt.dr.drugrecal

生产企业发布的召回信息，按照批次进行召回，收货和发货环节的单据处理中调用接口进行查询； */
type AlibabaAlihealthDrugKytDrDrugrecalAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel
}

// AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel is 疫苗药品召回 成功返回结果
type AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_drugrecal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrDrugrecalResult `json:"result,omitempty" xml:"result,omitempty"`
}
