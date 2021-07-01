package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundConfirmAPIResponse
确认资金单 API返回值
taobao.alitrip.axin.trans.fund.confirm

通过外部订单号进行资金结算 */
type TaobaoAlitripAxinTransFundConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransFundConfirmAPIResponseModel
}

// TaobaoAlitripAxinTransFundConfirmAPIResponseModel is 确认资金单 成功返回结果
type TaobaoAlitripAxinTransFundConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_fund_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 简单数据类型出参，用于测试top接入流程
	Result *TaobaoAlitripAxinTransFundConfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}
