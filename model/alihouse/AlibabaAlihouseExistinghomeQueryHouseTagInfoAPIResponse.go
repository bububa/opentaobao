package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse 活动标查询接口 API返回值
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
type AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponseModel is 活动标查询接口 成功返回结果
type AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_query_house_tag_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAlihouseExistinghomeQueryHouseTagInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse
func GetAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse() *AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse {
	return poolAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse.Get().(*AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse 将 AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse(v *AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse.Put(v)
}
