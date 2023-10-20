package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse 门店等级评分查询 API返回值
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
type AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel is 门店等级评分查询 成功返回结果
type AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_level_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlihouseExistinghomeStoreLevelQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse
func GetAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse() *AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse {
	return poolAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse.Get().(*AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse 将 AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse(v *AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse.Put(v)
}
