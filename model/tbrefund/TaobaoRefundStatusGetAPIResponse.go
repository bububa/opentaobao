package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundstatusgetAPIResponse 查询退款状态 API返回值
// taobao.refund.status.get
//
// 根据主订单或者子订单id查询退款状态,入参传入主订单或者子订单号1.如果当传入子订单时，返回子订单最后一笔退款单的状态,如果子订单申请退款退款返回空list.2.如果传传入主订单，则返回主订单下所有子订单的最后一笔退款单状态，如果对应的子订单没有生成退款单，则对应子订单对应数据返回。
type TaobaorefundstatusgetAPIResponse struct {
	model.CommonResponse
	TaobaorefundstatusgetAPIResponseModel
}

// TaobaorefundstatusgetAPIResponseModel is 查询退款状态 成功返回结果
type TaobaorefundstatusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	ResultPackage *TaobaorefundstatusgetResultSet `json:"result_package,omitempty" xml:"result_package,omitempty"`
}
