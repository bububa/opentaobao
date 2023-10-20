package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse 委托房源变更标准房源 API返回值
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
type AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel is 委托房源变更标准房源 成功返回结果
type AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_change_standard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseChangeStandardResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse
func GetAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse() *AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse 将 AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse(v *AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse.Put(v)
}
