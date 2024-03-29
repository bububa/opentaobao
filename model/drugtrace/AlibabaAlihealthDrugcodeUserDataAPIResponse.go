package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeUserDataAPIResponse 西安杨森同步用户行为接口 API返回值
// alibaba.alihealth.drugcode.user.data
//
// 西安杨森同步用户行为接口
type AlibabaAlihealthDrugcodeUserDataAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeUserDataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeUserDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeUserDataAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeUserDataAPIResponseModel is 西安杨森同步用户行为接口 成功返回结果
type AlibabaAlihealthDrugcodeUserDataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_user_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugcodeUserDataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeUserDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugcodeUserDataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeUserDataAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeUserDataAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeUserDataAPIResponse
func GetAlibabaAlihealthDrugcodeUserDataAPIResponse() *AlibabaAlihealthDrugcodeUserDataAPIResponse {
	return poolAlibabaAlihealthDrugcodeUserDataAPIResponse.Get().(*AlibabaAlihealthDrugcodeUserDataAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeUserDataAPIResponse 将 AlibabaAlihealthDrugcodeUserDataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeUserDataAPIResponse(v *AlibabaAlihealthDrugcodeUserDataAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeUserDataAPIResponse.Put(v)
}
