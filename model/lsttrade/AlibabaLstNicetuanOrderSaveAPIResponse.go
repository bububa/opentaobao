package lsttrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstNicetuanOrderSaveAPIResponse 十荟团订单同步至零售通 API返回值
// alibaba.lst.nicetuan.order.save
//
// 十荟团订单同步至零售通，十荟团单向写到零售通
type AlibabaLstNicetuanOrderSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLstNicetuanOrderSaveAPIResponseModel
}

// AlibabaLstNicetuanOrderSaveAPIResponseModel is 十荟团订单同步至零售通 成功返回结果
type AlibabaLstNicetuanOrderSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_nicetuan_order_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HsfResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
