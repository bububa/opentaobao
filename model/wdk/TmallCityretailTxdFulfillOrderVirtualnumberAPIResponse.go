package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse 淘鲜达虚拟号服务接口 API返回值
// tmall.cityretail.txd.fulfill.order.virtualnumber
//
// 虚拟小号绑定接口，只有开通了虚拟号服务的淘鲜达商家能绑定。
type TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse struct {
	model.CommonResponse
	TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel).Reset()
}

// TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel is 淘鲜达虚拟号服务接口 成功返回结果
type TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_txd_fulfill_order_virtualnumber_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse)
	},
}

// GetTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse 从 sync.Pool 获取 TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse
func GetTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse() *TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse {
	return poolTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse.Get().(*TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse)
}

// ReleaseTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse 将 TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse 保存到 sync.Pool
func ReleaseTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse(v *TmallCityretailTxdFulfillOrderVirtualnumberAPIResponse) {
	v.Reset()
	poolTmallCityretailTxdFulfillOrderVirtualnumberAPIResponse.Put(v)
}
