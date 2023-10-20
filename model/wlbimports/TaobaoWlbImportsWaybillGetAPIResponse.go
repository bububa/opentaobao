package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsWaybillGetAPIResponse 获取运单信息【云打印】 API返回值
// taobao.wlb.imports.waybill.get
//
// 一般进口商家，获取订单的电子面单链接地址
type TaobaoWlbImportsWaybillGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportsWaybillGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportsWaybillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportsWaybillGetAPIResponseModel).Reset()
}

// TaobaoWlbImportsWaybillGetAPIResponseModel is 获取运单信息【云打印】 成功返回结果
type TaobaoWlbImportsWaybillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_waybill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 电子面单链接地址
	Waybillurl string `json:"waybillurl,omitempty" xml:"waybillurl,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportsWaybillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Waybillurl = ""
}

var poolTaobaoWlbImportsWaybillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsWaybillGetAPIResponse)
	},
}

// GetTaobaoWlbImportsWaybillGetAPIResponse 从 sync.Pool 获取 TaobaoWlbImportsWaybillGetAPIResponse
func GetTaobaoWlbImportsWaybillGetAPIResponse() *TaobaoWlbImportsWaybillGetAPIResponse {
	return poolTaobaoWlbImportsWaybillGetAPIResponse.Get().(*TaobaoWlbImportsWaybillGetAPIResponse)
}

// ReleaseTaobaoWlbImportsWaybillGetAPIResponse 将 TaobaoWlbImportsWaybillGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportsWaybillGetAPIResponse(v *TaobaoWlbImportsWaybillGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportsWaybillGetAPIResponse.Put(v)
}
