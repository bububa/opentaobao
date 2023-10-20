package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse 租房合作品牌更新接口 API返回值
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
type AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel is 租房合作品牌更新接口 成功返回结果
type AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_cooperate_brand_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值对象
	Result *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse
func GetAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse() *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse {
	return poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse.Get().(*AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse 将 AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse(v *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse.Put(v)
}
