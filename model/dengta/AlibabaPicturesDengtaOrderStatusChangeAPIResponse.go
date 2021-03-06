package dengta

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaOrderStatusChangeAPIResponse 天下秀订单状态变更通知 API返回值
// alibaba.pictures.dengta.order.status.change
//
// 天下秀订单状态变更通知
type AlibabaPicturesDengtaOrderStatusChangeAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaOrderStatusChangeAPIResponseModel
}

// AlibabaPicturesDengtaOrderStatusChangeAPIResponseModel is 天下秀订单状态变更通知 成功返回结果
type AlibabaPicturesDengtaOrderStatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_order_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}
