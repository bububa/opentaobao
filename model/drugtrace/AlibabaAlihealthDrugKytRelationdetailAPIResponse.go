package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytRelationdetailAPIResponse
关联关系处理详情 API返回值
alibaba.alihealth.drug.kyt.relationdetail

关联关系处理详情 */
type AlibabaAlihealthDrugKytRelationdetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytRelationdetailAPIResponseModel
}

// AlibabaAlihealthDrugKytRelationdetailAPIResponseModel is 关联关系处理详情 成功返回结果
type AlibabaAlihealthDrugKytRelationdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_relationdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytRelationdetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
