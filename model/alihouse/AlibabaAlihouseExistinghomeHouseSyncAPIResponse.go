package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseSyncAPIResponse 房源信息同步 API返回值
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
type AlibabaAlihouseExistinghomeHouseSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel is 房源信息同步 成功返回结果
type AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseSyncAPIResponse
func GetAlibabaAlihouseExistinghomeHouseSyncAPIResponse() *AlibabaAlihouseExistinghomeHouseSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseSyncAPIResponse 将 AlibabaAlihouseExistinghomeHouseSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseSyncAPIResponse(v *AlibabaAlihouseExistinghomeHouseSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseSyncAPIResponse.Put(v)
}
