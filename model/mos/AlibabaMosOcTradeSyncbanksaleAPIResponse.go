package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosoctradesyncbanksaleAPIResponse 云闪付、银行卡销售数据回传接口 API返回值
// alibaba.mos.oc.trade.syncbanksale
//
// 云闪付、银行卡销售数据回传
type AlibabamosoctradesyncbanksaleAPIResponse struct {
	model.CommonResponse
	AlibabamosoctradesyncbanksaleAPIResponseModel
}

// AlibabamosoctradesyncbanksaleAPIResponseModel is 云闪付、银行卡销售数据回传接口 成功返回结果
type AlibabamosoctradesyncbanksaleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_oc_trade_syncbanksale_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultDTO
	Result *AlibabamosoctradesyncbanksaleResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
