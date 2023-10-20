package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillShengxianGetAPIResponse 商家获取生鲜电子面单号 API返回值
// taobao.wlb.waybill.shengxian.get
//
// 商家通过交易订单号获取电子面单接口
type TaobaoWlbWaybillShengxianGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillShengxianGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillShengxianGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWaybillShengxianGetAPIResponseModel).Reset()
}

// TaobaoWlbWaybillShengxianGetAPIResponseModel is 商家获取生鲜电子面单号 成功返回结果
type TaobaoWlbWaybillShengxianGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_shengxian_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功后返回的生鲜电子面单信息
	FreshWaybill *FreshWaybill `json:"fresh_waybill,omitempty" xml:"fresh_waybill,omitempty"`
	// 生成是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillShengxianGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FreshWaybill = nil
	m.IsSuccess = false
}

var poolTaobaoWlbWaybillShengxianGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWaybillShengxianGetAPIResponse)
	},
}

// GetTaobaoWlbWaybillShengxianGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWaybillShengxianGetAPIResponse
func GetTaobaoWlbWaybillShengxianGetAPIResponse() *TaobaoWlbWaybillShengxianGetAPIResponse {
	return poolTaobaoWlbWaybillShengxianGetAPIResponse.Get().(*TaobaoWlbWaybillShengxianGetAPIResponse)
}

// ReleaseTaobaoWlbWaybillShengxianGetAPIResponse 将 TaobaoWlbWaybillShengxianGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWaybillShengxianGetAPIResponse(v *TaobaoWlbWaybillShengxianGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWaybillShengxianGetAPIResponse.Put(v)
}
