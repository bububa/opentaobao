package wdk

import (
	"encoding/xml"

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

// TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel is 淘鲜达虚拟号服务接口 成功返回结果
type TmallCityretailTxdFulfillOrderVirtualnumberAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_txd_fulfill_order_virtualnumber_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}
