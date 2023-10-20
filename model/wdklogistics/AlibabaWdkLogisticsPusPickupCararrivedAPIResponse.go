package wdklogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkLogisticsPusPickupCararrivedAPIResponse 自提业务-车辆到达上报车牌号 API返回值
// alibaba.wdk.logistics.pus.pickup.cararrived
//
// 自提业务-汽车自提,车辆到达上报车牌号
type AlibabaWdkLogisticsPusPickupCararrivedAPIResponse struct {
	model.CommonResponse
	AlibabaWdkLogisticsPusPickupCararrivedAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkLogisticsPusPickupCararrivedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkLogisticsPusPickupCararrivedAPIResponseModel).Reset()
}

// AlibabaWdkLogisticsPusPickupCararrivedAPIResponseModel is 自提业务-车辆到达上报车牌号 成功返回结果
type AlibabaWdkLogisticsPusPickupCararrivedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_logistics_pus_pickup_cararrived_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result根结点
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkLogisticsPusPickupCararrivedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkLogisticsPusPickupCararrivedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkLogisticsPusPickupCararrivedAPIResponse)
	},
}

// GetAlibabaWdkLogisticsPusPickupCararrivedAPIResponse 从 sync.Pool 获取 AlibabaWdkLogisticsPusPickupCararrivedAPIResponse
func GetAlibabaWdkLogisticsPusPickupCararrivedAPIResponse() *AlibabaWdkLogisticsPusPickupCararrivedAPIResponse {
	return poolAlibabaWdkLogisticsPusPickupCararrivedAPIResponse.Get().(*AlibabaWdkLogisticsPusPickupCararrivedAPIResponse)
}

// ReleaseAlibabaWdkLogisticsPusPickupCararrivedAPIResponse 将 AlibabaWdkLogisticsPusPickupCararrivedAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkLogisticsPusPickupCararrivedAPIResponse(v *AlibabaWdkLogisticsPusPickupCararrivedAPIResponse) {
	v.Reset()
	poolAlibabaWdkLogisticsPusPickupCararrivedAPIResponse.Put(v)
}
