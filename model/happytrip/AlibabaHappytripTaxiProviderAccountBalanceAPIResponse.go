package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiprovideraccountbalanceAPIResponse 供应商渠道余额 API返回值
// alibaba.happytrip.taxi.provider.account.balance
//
// 查询不同供应商不同渠道账户余额
type AlibabahappytriptaxiprovideraccountbalanceAPIResponse struct {
	model.CommonResponse
	AlibabahappytriptaxiprovideraccountbalanceAPIResponseModel
}

// AlibabahappytriptaxiprovideraccountbalanceAPIResponseModel is 供应商渠道余额 成功返回结果
type AlibabahappytriptaxiprovideraccountbalanceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_provider_account_balance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 数据
	Data *AlibabahappytriptaxiprovideraccountbalanceData `json:"data,omitempty" xml:"data,omitempty"`
}
