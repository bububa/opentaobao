package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseUpselfAPIResponse 房源上架 API返回值
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
type AlibabaAlihouseExistinghomeHouseUpselfAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseUpselfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel is 房源上架 成功返回结果
type AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_upself_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseUpselfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseUpselfAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseUpselfAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseUpselfAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseUpselfAPIResponse
func GetAlibabaAlihouseExistinghomeHouseUpselfAPIResponse() *AlibabaAlihouseExistinghomeHouseUpselfAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseUpselfAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseUpselfAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseUpselfAPIResponse 将 AlibabaAlihouseExistinghomeHouseUpselfAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseUpselfAPIResponse(v *AlibabaAlihouseExistinghomeHouseUpselfAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseUpselfAPIResponse.Put(v)
}
