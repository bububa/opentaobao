package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderPrintSaleJudgeAPIResponse
导购员判断 API返回值
taobao.omniorder.print.sale.judge

用于判断当前子账号是否导购员 */
type TaobaoOmniorderPrintSaleJudgeAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderPrintSaleJudgeAPIResponseModel
}

// TaobaoOmniorderPrintSaleJudgeAPIResponseModel is 导购员判断 成功返回结果
type TaobaoOmniorderPrintSaleJudgeAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_print_sale_judge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
