package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotraderestpayfeemodifyAPIResponse 分阶段订单尾款改价 API返回值
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
type TmallaliautotraderestpayfeemodifyAPIResponse struct {
	model.CommonResponse
	TmallaliautotraderestpayfeemodifyAPIResponseModel
}

// TmallaliautotraderestpayfeemodifyAPIResponseModel is 分阶段订单尾款改价 成功返回结果
type TmallaliautotraderestpayfeemodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_restpayfee_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
