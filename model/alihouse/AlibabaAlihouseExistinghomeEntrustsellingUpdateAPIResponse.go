package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse 管家状态及房源信息接口 API返回值
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
type AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponseModel is 管家状态及房源信息接口 成功返回结果
type AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_entrustselling_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeEntrustsellingUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse
func GetAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse() *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse {
	return poolAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse.Get().(*AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse 将 AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse(v *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse.Put(v)
}
