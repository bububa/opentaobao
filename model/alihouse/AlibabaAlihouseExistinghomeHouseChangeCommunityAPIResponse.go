package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse 房屋、房源变更所属小区 API返回值
// alibaba.alihouse.existinghome.house.change.community
//
// 房屋、房源变更所属小区
type AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponseModel is 房屋、房源变更所属小区 成功返回结果
type AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_change_community_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihouseExistinghomeHouseChangeCommunityResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse
func GetAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse() *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse 将 AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse(v *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse.Put(v)
}
