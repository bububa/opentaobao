package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQuerybatchprodAPIResponse 批次产品查询(根据企业名和批次号查询产品信息) API返回值
// alibaba.alihealth.drug.kyt.querybatchprod
//
// 根据企业名和批次号查询药品信息，支持使用更名之前的老企业名查询，支持批次号大小写模糊，应用于药店或医院入库环节，通过在入库环节获取赋码的产品目录，可强制要求对相应的产品必须进行扫码入库；
type AlibabaAlihealthDrugKytQuerybatchprodAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQuerybatchprodAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQuerybatchprodAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQuerybatchprodAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQuerybatchprodAPIResponseModel is 批次产品查询(根据企业名和批次号查询产品信息) 成功返回结果
type AlibabaAlihealthDrugKytQuerybatchprodAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_querybatchprod_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaAlihealthDrugKytQuerybatchprodResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQuerybatchprodAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQuerybatchprodAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQuerybatchprodAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQuerybatchprodAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQuerybatchprodAPIResponse
func GetAlibabaAlihealthDrugKytQuerybatchprodAPIResponse() *AlibabaAlihealthDrugKytQuerybatchprodAPIResponse {
	return poolAlibabaAlihealthDrugKytQuerybatchprodAPIResponse.Get().(*AlibabaAlihealthDrugKytQuerybatchprodAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQuerybatchprodAPIResponse 将 AlibabaAlihealthDrugKytQuerybatchprodAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQuerybatchprodAPIResponse(v *AlibabaAlihealthDrugKytQuerybatchprodAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQuerybatchprodAPIResponse.Put(v)
}
