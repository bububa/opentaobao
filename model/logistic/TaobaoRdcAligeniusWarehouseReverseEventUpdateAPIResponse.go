package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse
销退单事件回传接口 API返回值
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台 */
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponseModel
}

// TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponseModel is 销退单事件回传接口 成功返回结果
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_warehouse_reverse_event_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoRdcAligeniusWarehouseReverseEventUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
