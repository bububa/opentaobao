package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangInventoryFutureplanAPIResponse 负卖计划 API返回值
// alibaba.dchain.aoxiang.inventory.futureplan
//
// 负卖计划。底层有白名单控制，并非对所有商家开放。如果需要使用，请联系对应的小二增加白名单
type AlibabaDchainAoxiangInventoryFutureplanAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangInventoryFutureplanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangInventoryFutureplanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangInventoryFutureplanAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangInventoryFutureplanAPIResponseModel is 负卖计划 成功返回结果
type AlibabaDchainAoxiangInventoryFutureplanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_inventory_futureplan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	FuturePlanResponse *TopResponse `json:"future_plan_response,omitempty" xml:"future_plan_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangInventoryFutureplanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FuturePlanResponse = nil
}

var poolAlibabaDchainAoxiangInventoryFutureplanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangInventoryFutureplanAPIResponse)
	},
}

// GetAlibabaDchainAoxiangInventoryFutureplanAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangInventoryFutureplanAPIResponse
func GetAlibabaDchainAoxiangInventoryFutureplanAPIResponse() *AlibabaDchainAoxiangInventoryFutureplanAPIResponse {
	return poolAlibabaDchainAoxiangInventoryFutureplanAPIResponse.Get().(*AlibabaDchainAoxiangInventoryFutureplanAPIResponse)
}

// ReleaseAlibabaDchainAoxiangInventoryFutureplanAPIResponse 将 AlibabaDchainAoxiangInventoryFutureplanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangInventoryFutureplanAPIResponse(v *AlibabaDchainAoxiangInventoryFutureplanAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangInventoryFutureplanAPIResponse.Put(v)
}
