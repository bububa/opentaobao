package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse 医院根据码查询码信息 API返回值
// alibaba.alihealth.drug.code.kyt.yy.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponseModel is 医院根据码查询码信息 成功返回结果
type AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_yy_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytYyQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse
func GetAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse() *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse 将 AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse(v *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytYyQuerycodeAPIResponse.Put(v)
}
