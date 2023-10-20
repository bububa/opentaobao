package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletchargeruleAPIResponse 流量钱包直充（根据号码归属地省份路由） API返回值
// alibaba.aliqin.flow.wallet.charge.rule
//
// 流量钱包直充（根据号码归属地省份路由）
type AlibabaaliqinflowwalletchargeruleAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinflowwalletchargeruleAPIResponseModel
}

// AlibabaaliqinflowwalletchargeruleAPIResponseModel is 流量钱包直充（根据号码归属地省份路由） 成功返回结果
type AlibabaaliqinflowwalletchargeruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_charge_rule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// {&#34;error&#34;:&#34;true&#34;,&#34;msg&#34;:&#34;返回信息&#34;}
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
}
