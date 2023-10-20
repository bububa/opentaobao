package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse 删除房源 API返回值
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
type AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponseModel is 删除房源 成功返回结果
type AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_delete_house_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseDeleteHouseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse
func GetAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse() *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse 将 AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse(v *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse.Put(v)
}
