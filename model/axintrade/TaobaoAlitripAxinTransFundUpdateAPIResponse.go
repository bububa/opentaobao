package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundUpdateAPIResponse
修改资金单接口 API返回值
taobao.alitrip.axin.trans.fund.update

阿信供销平台-修改资金单接口 */
type TaobaoAlitripAxinTransFundUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransFundUpdateAPIResponseModel
}

// TaobaoAlitripAxinTransFundUpdateAPIResponseModel is 修改资金单接口 成功返回结果
type TaobaoAlitripAxinTransFundUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_fund_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripAxinTransFundUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
