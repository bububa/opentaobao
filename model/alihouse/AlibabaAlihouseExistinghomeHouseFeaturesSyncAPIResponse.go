package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse 房源标准打标数据同步 API返回值
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
type AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel is 房源标准打标数据同步 成功返回结果
type AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_features_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseFeaturesSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse
func GetAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse() *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse 将 AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse(v *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse.Put(v)
}
