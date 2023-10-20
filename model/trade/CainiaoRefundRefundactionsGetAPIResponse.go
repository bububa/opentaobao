package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaorefundrefundactionsgetAPIResponse 判断该订单能执行的逆向操作集合列表 API返回值
// cainiao.refund.refundactions.get
//
// 判断该订单能执行的逆向操作集合列表
type CainiaorefundrefundactionsgetAPIResponse struct {
	model.CommonResponse
	CainiaorefundrefundactionsgetAPIResponseModel
}

// CainiaorefundrefundactionsgetAPIResponseModel is 判断该订单能执行的逆向操作集合列表 成功返回结果
type CainiaorefundrefundactionsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_refund_refundactions_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *CainiaorefundrefundactionsgetBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
