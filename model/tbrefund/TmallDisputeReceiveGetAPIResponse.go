package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalldisputereceivegetAPIResponse 天猫逆向纠纷查询 API返回值
// tmall.dispute.receive.get
//
// 展示商家所有退款信息
type TmalldisputereceivegetAPIResponse struct {
	model.CommonResponse
	TmalldisputereceivegetAPIResponseModel
}

// TmalldisputereceivegetAPIResponseModel is 天猫逆向纠纷查询 成功返回结果
type TmalldisputereceivegetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_dispute_receive_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmalldisputereceivegetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
