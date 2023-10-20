package lstpos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopencashiersynccashierdataAPIResponse 收银快照同步接口(最多10条订单信息) API返回值
// alibaba.lst.pos.open.cashier.synccashierdata
//
// 收银快照同步接口(最多10条订单信息)
type AlibabalstposopencashiersynccashierdataAPIResponse struct {
	model.CommonResponse
	AlibabalstposopencashiersynccashierdataAPIResponseModel
}

// AlibabalstposopencashiersynccashierdataAPIResponseModel is 收银快照同步接口(最多10条订单信息) 成功返回结果
type AlibabalstposopencashiersynccashierdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_cashier_synccashierdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *AlibabalstposopencashiersynccashierdataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
