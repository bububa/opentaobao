package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniuswarehouseresendupdateAPIResponse 补发单状态回传 API返回值
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
type TaobaordcaligeniuswarehouseresendupdateAPIResponse struct {
	model.CommonResponse
	TaobaordcaligeniuswarehouseresendupdateAPIResponseModel
}

// TaobaordcaligeniuswarehouseresendupdateAPIResponseModel is 补发单状态回传 成功返回结果
type TaobaordcaligeniuswarehouseresendupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_warehouse_resend_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaordcaligeniuswarehouseresendupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
