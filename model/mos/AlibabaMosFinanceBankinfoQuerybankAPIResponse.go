package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosfinancebankinfoquerybankAPIResponse 供应商银行账号查询 API返回值
// alibaba.mos.finance.bankinfo.querybank
//
// 查询供应商对应的银行账号信息
type AlibabamosfinancebankinfoquerybankAPIResponse struct {
	model.CommonResponse
	AlibabamosfinancebankinfoquerybankAPIResponseModel
}

// AlibabamosfinancebankinfoquerybankAPIResponseModel is 供应商银行账号查询 成功返回结果
type AlibabamosfinancebankinfoquerybankAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_finance_bankinfo_querybank_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamosfinancebankinfoquerybankResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
