package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarfpcarrestpayreceiveAPIResponse 门店线下已收尾款 API返回值
// tmall.car.fpcar.restpay.receive
//
// 提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
type TmallcarfpcarrestpayreceiveAPIResponse struct {
	model.CommonResponse
	TmallcarfpcarrestpayreceiveAPIResponseModel
}

// TmallcarfpcarrestpayreceiveAPIResponseModel is 门店线下已收尾款 成功返回结果
type TmallcarfpcarrestpayreceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_fpcar_restpay_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 业务结果信息详细描述
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// 是否成功
	Succes bool `json:"succes,omitempty" xml:"succes,omitempty"`
}
