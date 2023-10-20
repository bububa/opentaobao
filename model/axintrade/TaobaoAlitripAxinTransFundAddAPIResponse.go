package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintransfundaddAPIResponse 创建资金单接口 API返回值
// taobao.alitrip.axin.trans.fund.add
//
// 创建资金单
type TaobaoalitripaxintransfundaddAPIResponse struct {
	model.CommonResponse
	TaobaoalitripaxintransfundaddAPIResponseModel
}

// TaobaoalitripaxintransfundaddAPIResponseModel is 创建资金单接口 成功返回结果
type TaobaoalitripaxintransfundaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_fund_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitripaxintransfundaddResult `json:"result,omitempty" xml:"result,omitempty"`
}
