package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressordertmsupdateAPIResponse 服务商修改上门取退时间接口 API返回值
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
type TaobaologisticsexpressordertmsupdateAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpressordertmsupdateAPIResponseModel
}

// TaobaologisticsexpressordertmsupdateAPIResponseModel is 服务商修改上门取退时间接口 成功返回结果
type TaobaologisticsexpressordertmsupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_order_tms_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 响应码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 系统成功失败
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
