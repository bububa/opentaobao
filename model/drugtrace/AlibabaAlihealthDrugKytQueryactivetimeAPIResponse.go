package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryactivetimeAPIResponse 药品激活状态同步 API返回值
// alibaba.alihealth.drug.kyt.queryactivetime
//
// 根据赋码资源（CodeVersion + resCode）获得最新激活时间
// 应用于各地市对接前进行药品目录匹配，医保中心存在的药品可能比较陈旧杂乱
type AlibabaAlihealthDrugKytQueryactivetimeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQueryactivetimeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryactivetimeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQueryactivetimeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQueryactivetimeAPIResponseModel is 药品激活状态同步 成功返回结果
type AlibabaAlihealthDrugKytQueryactivetimeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_queryactivetime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytQueryactivetimeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryactivetimeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQueryactivetimeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryactivetimeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQueryactivetimeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryactivetimeAPIResponse
func GetAlibabaAlihealthDrugKytQueryactivetimeAPIResponse() *AlibabaAlihealthDrugKytQueryactivetimeAPIResponse {
	return poolAlibabaAlihealthDrugKytQueryactivetimeAPIResponse.Get().(*AlibabaAlihealthDrugKytQueryactivetimeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQueryactivetimeAPIResponse 将 AlibabaAlihealthDrugKytQueryactivetimeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryactivetimeAPIResponse(v *AlibabaAlihealthDrugKytQueryactivetimeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryactivetimeAPIResponse.Put(v)
}
