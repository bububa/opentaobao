package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaobmsorderconsignconfirmAPIResponse BMS出库通知 API返回值
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
type CainiaobmsorderconsignconfirmAPIResponse struct {
	model.CommonResponse
	CainiaobmsorderconsignconfirmAPIResponseModel
}

// CainiaobmsorderconsignconfirmAPIResponseModel is BMS出库通知 成功返回结果
type CainiaobmsorderconsignconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_bms_order_consign_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
