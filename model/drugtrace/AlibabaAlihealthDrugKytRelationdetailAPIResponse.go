package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytRelationdetailAPIResponse 关联关系处理详情 API返回值
// alibaba.alihealth.drug.kyt.relationdetail
//
// 关联关系处理详情
type AlibabaAlihealthDrugKytRelationdetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytRelationdetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytRelationdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytRelationdetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytRelationdetailAPIResponseModel is 关联关系处理详情 成功返回结果
type AlibabaAlihealthDrugKytRelationdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_relationdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytRelationdetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytRelationdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytRelationdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytRelationdetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytRelationdetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytRelationdetailAPIResponse
func GetAlibabaAlihealthDrugKytRelationdetailAPIResponse() *AlibabaAlihealthDrugKytRelationdetailAPIResponse {
	return poolAlibabaAlihealthDrugKytRelationdetailAPIResponse.Get().(*AlibabaAlihealthDrugKytRelationdetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytRelationdetailAPIResponse 将 AlibabaAlihealthDrugKytRelationdetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytRelationdetailAPIResponse(v *AlibabaAlihealthDrugKytRelationdetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytRelationdetailAPIResponse.Put(v)
}
