package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse 经纪人积分同步 API返回值
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
type AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel is 经纪人积分同步 成功返回结果
type AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_broker_points_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBrokerPointsSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse
func GetAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse() *AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse 将 AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse(v *AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse.Put(v)
}
