package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWlborderGetAPIResponse 根据物流宝订单编号查询物流宝订单概要信息 API返回值
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
type TaobaoWlbWlborderGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWlborderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWlborderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWlborderGetAPIResponseModel).Reset()
}

// TaobaoWlbWlborderGetAPIResponseModel is 根据物流宝订单编号查询物流宝订单概要信息 成功返回结果
type TaobaoWlbWlborderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wlborder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流宝订单对象
	WlbOrder *WlbOrder `json:"wlb_order,omitempty" xml:"wlb_order,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWlborderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WlbOrder = nil
}

var poolTaobaoWlbWlborderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWlborderGetAPIResponse)
	},
}

// GetTaobaoWlbWlborderGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWlborderGetAPIResponse
func GetTaobaoWlbWlborderGetAPIResponse() *TaobaoWlbWlborderGetAPIResponse {
	return poolTaobaoWlbWlborderGetAPIResponse.Get().(*TaobaoWlbWlborderGetAPIResponse)
}

// ReleaseTaobaoWlbWlborderGetAPIResponse 将 TaobaoWlbWlborderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWlborderGetAPIResponse(v *TaobaoWlbWlborderGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWlborderGetAPIResponse.Put(v)
}
