package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityStoreGetstorelistAPIResponse ISV查询门店 API返回值
// alibaba.lsy.crm.activity.store.getstorelist
//
// ISV查询门店
type AlibabaLsyCrmActivityStoreGetstorelistAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityStoreGetstorelistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityStoreGetstorelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivityStoreGetstorelistAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivityStoreGetstorelistAPIResponseModel is ISV查询门店 成功返回结果
type AlibabaLsyCrmActivityStoreGetstorelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_store_getstorelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	PageResultDO *PageResultDo `json:"page_result_d_o,omitempty" xml:"page_result_d_o,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityStoreGetstorelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PageResultDO = nil
}

var poolAlibabaLsyCrmActivityStoreGetstorelistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityStoreGetstorelistAPIResponse)
	},
}

// GetAlibabaLsyCrmActivityStoreGetstorelistAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivityStoreGetstorelistAPIResponse
func GetAlibabaLsyCrmActivityStoreGetstorelistAPIResponse() *AlibabaLsyCrmActivityStoreGetstorelistAPIResponse {
	return poolAlibabaLsyCrmActivityStoreGetstorelistAPIResponse.Get().(*AlibabaLsyCrmActivityStoreGetstorelistAPIResponse)
}

// ReleaseAlibabaLsyCrmActivityStoreGetstorelistAPIResponse 将 AlibabaLsyCrmActivityStoreGetstorelistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivityStoreGetstorelistAPIResponse(v *AlibabaLsyCrmActivityStoreGetstorelistAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivityStoreGetstorelistAPIResponse.Put(v)
}
