package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytStorebilllistAPIResponse 零售端平台单据查询 API返回值
// alibaba.alihealth.drug.kyt.storebilllist
//
// 零售端平台单据查询
type AlibabaAlihealthDrugKytStorebilllistAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytStorebilllistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytStorebilllistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytStorebilllistAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytStorebilllistAPIResponseModel is 零售端平台单据查询 成功返回结果
type AlibabaAlihealthDrugKytStorebilllistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_storebilllist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytStorebilllistResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytStorebilllistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytStorebilllistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytStorebilllistAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytStorebilllistAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytStorebilllistAPIResponse
func GetAlibabaAlihealthDrugKytStorebilllistAPIResponse() *AlibabaAlihealthDrugKytStorebilllistAPIResponse {
	return poolAlibabaAlihealthDrugKytStorebilllistAPIResponse.Get().(*AlibabaAlihealthDrugKytStorebilllistAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytStorebilllistAPIResponse 将 AlibabaAlihealthDrugKytStorebilllistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytStorebilllistAPIResponse(v *AlibabaAlihealthDrugKytStorebilllistAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytStorebilllistAPIResponse.Put(v)
}
