package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghouseHouseBaseAPIResponse （租房）同步房屋信息 API返回值
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
type AlibabaAlihouseExistinghouseHouseBaseAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghouseHouseBaseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghouseHouseBaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghouseHouseBaseAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghouseHouseBaseAPIResponseModel is （租房）同步房屋信息 成功返回结果
type AlibabaAlihouseExistinghouseHouseBaseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghouse_house_base_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseExistinghouseHouseBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghouseHouseBaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghouseHouseBaseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghouseHouseBaseAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghouseHouseBaseAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghouseHouseBaseAPIResponse
func GetAlibabaAlihouseExistinghouseHouseBaseAPIResponse() *AlibabaAlihouseExistinghouseHouseBaseAPIResponse {
	return poolAlibabaAlihouseExistinghouseHouseBaseAPIResponse.Get().(*AlibabaAlihouseExistinghouseHouseBaseAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghouseHouseBaseAPIResponse 将 AlibabaAlihouseExistinghouseHouseBaseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghouseHouseBaseAPIResponse(v *AlibabaAlihouseExistinghouseHouseBaseAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghouseHouseBaseAPIResponse.Put(v)
}
