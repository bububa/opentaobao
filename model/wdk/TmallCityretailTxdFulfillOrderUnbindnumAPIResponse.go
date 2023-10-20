package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcityretailtxdfulfillorderunbindnumAPIResponse 淘鲜达虚拟号服务解绑接口 API返回值
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
type TmallcityretailtxdfulfillorderunbindnumAPIResponse struct {
	model.CommonResponse
	TmallcityretailtxdfulfillorderunbindnumAPIResponseModel
}

// TmallcityretailtxdfulfillorderunbindnumAPIResponseModel is 淘鲜达虚拟号服务解绑接口 成功返回结果
type TmallcityretailtxdfulfillorderunbindnumAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_txd_fulfill_order_unbindnum_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}
