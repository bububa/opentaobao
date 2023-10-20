package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixdistributioncmbparamencryptAPIResponse 加密招商一网能支付入参 API返回值
// alibaba.damai.maitix.distribution.cmb.paramencrypt
//
// encryptParam4Cmb
type AlibabadamaimaitixdistributioncmbparamencryptAPIResponse struct {
	model.CommonResponse
	AlibabadamaimaitixdistributioncmbparamencryptAPIResponseModel
}

// AlibabadamaimaitixdistributioncmbparamencryptAPIResponseModel is 加密招商一网能支付入参 成功返回结果
type AlibabadamaimaitixdistributioncmbparamencryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_cmb_paramencrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
