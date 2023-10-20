package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoBmsOrderConsignConfirmAPIResponse BMS出库通知 API返回值
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
type CainiaoBmsOrderConsignConfirmAPIResponse struct {
	model.CommonResponse
	CainiaoBmsOrderConsignConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoBmsOrderConsignConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoBmsOrderConsignConfirmAPIResponseModel).Reset()
}

// CainiaoBmsOrderConsignConfirmAPIResponseModel is BMS出库通知 成功返回结果
type CainiaoBmsOrderConsignConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_bms_order_consign_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoBmsOrderConsignConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoBmsOrderConsignConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoBmsOrderConsignConfirmAPIResponse)
	},
}

// GetCainiaoBmsOrderConsignConfirmAPIResponse 从 sync.Pool 获取 CainiaoBmsOrderConsignConfirmAPIResponse
func GetCainiaoBmsOrderConsignConfirmAPIResponse() *CainiaoBmsOrderConsignConfirmAPIResponse {
	return poolCainiaoBmsOrderConsignConfirmAPIResponse.Get().(*CainiaoBmsOrderConsignConfirmAPIResponse)
}

// ReleaseCainiaoBmsOrderConsignConfirmAPIResponse 将 CainiaoBmsOrderConsignConfirmAPIResponse 保存到 sync.Pool
func ReleaseCainiaoBmsOrderConsignConfirmAPIResponse(v *CainiaoBmsOrderConsignConfirmAPIResponse) {
	v.Reset()
	poolCainiaoBmsOrderConsignConfirmAPIResponse.Put(v)
}
