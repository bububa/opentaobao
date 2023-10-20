package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwuspbillreordaddAPIResponse 内购服务确认单明细上传接口 API返回值
// taobao.fuwu.sp.billreord.add
//
// isv能通过该接口上传确认单明细数据
type TaobaofuwuspbillreordaddAPIResponse struct {
	model.CommonResponse
	TaobaofuwuspbillreordaddAPIResponseModel
}

// TaobaofuwuspbillreordaddAPIResponseModel is 内购服务确认单明细上传接口 成功返回结果
type TaobaofuwuspbillreordaddAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_sp_billreord_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回调用结果
	AddResult bool `json:"add_result,omitempty" xml:"add_result,omitempty"`
}
