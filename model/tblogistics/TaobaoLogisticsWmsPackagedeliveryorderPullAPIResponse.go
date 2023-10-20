package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackagedeliveryorderpullAPIResponse 包裹出库单拉单 API返回值
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
type TaobaologisticswmspackagedeliveryorderpullAPIResponse struct {
	model.CommonResponse
	TaobaologisticswmspackagedeliveryorderpullAPIResponseModel
}

// TaobaologisticswmspackagedeliveryorderpullAPIResponseModel is 包裹出库单拉单 成功返回结果
type TaobaologisticswmspackagedeliveryorderpullAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packagedeliveryorder_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
