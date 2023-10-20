package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse 同城配送在线下单检查授权 API返回值
// alibaba.ascp.logistics.instantsonline.checkdeliveryauth
//
// 同城配送在线下单检查授权
type AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponseModel is 同城配送在线下单检查授权 成功返回结果
type AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_checkdeliveryauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse
func GetAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse() *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse {
	return poolAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse.Get().(*AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse 将 AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse(v *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse.Put(v)
}
