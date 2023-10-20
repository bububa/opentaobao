package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailTxdFulfillOrderUnbindnumAPIResponse 淘鲜达虚拟号服务解绑接口 API返回值
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
type TmallCityretailTxdFulfillOrderUnbindnumAPIResponse struct {
	model.CommonResponse
	TmallCityretailTxdFulfillOrderUnbindnumAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCityretailTxdFulfillOrderUnbindnumAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCityretailTxdFulfillOrderUnbindnumAPIResponseModel).Reset()
}

// TmallCityretailTxdFulfillOrderUnbindnumAPIResponseModel is 淘鲜达虚拟号服务解绑接口 成功返回结果
type TmallCityretailTxdFulfillOrderUnbindnumAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_txd_fulfill_order_unbindnum_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCityretailTxdFulfillOrderUnbindnumAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCityretailTxdFulfillOrderUnbindnumAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCityretailTxdFulfillOrderUnbindnumAPIResponse)
	},
}

// GetTmallCityretailTxdFulfillOrderUnbindnumAPIResponse 从 sync.Pool 获取 TmallCityretailTxdFulfillOrderUnbindnumAPIResponse
func GetTmallCityretailTxdFulfillOrderUnbindnumAPIResponse() *TmallCityretailTxdFulfillOrderUnbindnumAPIResponse {
	return poolTmallCityretailTxdFulfillOrderUnbindnumAPIResponse.Get().(*TmallCityretailTxdFulfillOrderUnbindnumAPIResponse)
}

// ReleaseTmallCityretailTxdFulfillOrderUnbindnumAPIResponse 将 TmallCityretailTxdFulfillOrderUnbindnumAPIResponse 保存到 sync.Pool
func ReleaseTmallCityretailTxdFulfillOrderUnbindnumAPIResponse(v *TmallCityretailTxdFulfillOrderUnbindnumAPIResponse) {
	v.Reset()
	poolTmallCityretailTxdFulfillOrderUnbindnumAPIResponse.Put(v)
}
