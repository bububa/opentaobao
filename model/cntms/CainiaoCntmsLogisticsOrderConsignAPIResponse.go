package cntms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntmsLogisticsOrderConsignAPIResponse 菜鸟配商家仓库发货 API返回值
// cainiao.cntms.logistics.order.consign
//
// 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaoCntmsLogisticsOrderConsignAPIResponse struct {
	model.CommonResponse
	CainiaoCntmsLogisticsOrderConsignAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCntmsLogisticsOrderConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCntmsLogisticsOrderConsignAPIResponseModel).Reset()
}

// CainiaoCntmsLogisticsOrderConsignAPIResponseModel is 菜鸟配商家仓库发货 成功返回结果
type CainiaoCntmsLogisticsOrderConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntms_logistics_order_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流单号
	LogisticsOrderCode string `json:"logistics_order_code,omitempty" xml:"logistics_order_code,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCntmsLogisticsOrderConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.LogisticsOrderCode = ""
}

var poolCainiaoCntmsLogisticsOrderConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCntmsLogisticsOrderConsignAPIResponse)
	},
}

// GetCainiaoCntmsLogisticsOrderConsignAPIResponse 从 sync.Pool 获取 CainiaoCntmsLogisticsOrderConsignAPIResponse
func GetCainiaoCntmsLogisticsOrderConsignAPIResponse() *CainiaoCntmsLogisticsOrderConsignAPIResponse {
	return poolCainiaoCntmsLogisticsOrderConsignAPIResponse.Get().(*CainiaoCntmsLogisticsOrderConsignAPIResponse)
}

// ReleaseCainiaoCntmsLogisticsOrderConsignAPIResponse 将 CainiaoCntmsLogisticsOrderConsignAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCntmsLogisticsOrderConsignAPIResponse(v *CainiaoCntmsLogisticsOrderConsignAPIResponse) {
	v.Reset()
	poolCainiaoCntmsLogisticsOrderConsignAPIResponse.Put(v)
}
