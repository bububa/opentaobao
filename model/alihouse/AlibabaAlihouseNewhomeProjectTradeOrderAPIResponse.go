package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojecttradeorderAPIResponse 旺铺楼盘和交易商品排序 API返回值
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
type AlibabaalihousenewhomeprojecttradeorderAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojecttradeorderAPIResponseModel
}

// AlibabaalihousenewhomeprojecttradeorderAPIResponseModel is 旺铺楼盘和交易商品排序 成功返回结果
type AlibabaalihousenewhomeprojecttradeorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_trade_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojecttradeorderResult `json:"result,omitempty" xml:"result,omitempty"`
}
