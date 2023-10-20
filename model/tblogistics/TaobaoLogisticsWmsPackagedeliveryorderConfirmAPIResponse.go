package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackagedeliveryorderconfirmAPIResponse 包裹出库单确认 API返回值
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
type TaobaologisticswmspackagedeliveryorderconfirmAPIResponse struct {
	model.CommonResponse
	TaobaologisticswmspackagedeliveryorderconfirmAPIResponseModel
}

// TaobaologisticswmspackagedeliveryorderconfirmAPIResponseModel is 包裹出库单确认 成功返回结果
type TaobaologisticswmspackagedeliveryorderconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packagedeliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
