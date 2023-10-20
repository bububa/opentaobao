package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse 加密招商一网能支付入参 API返回值
// alibaba.damai.maitix.distribution.cmb.paramencrypt
//
// encryptParam4Cmb
type AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponseModel is 加密招商一网能支付入参 成功返回结果
type AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_cmb_paramencrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse)
	},
}

// GetAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse
func GetAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse() *AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse {
	return poolAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse.Get().(*AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse)
}

// ReleaseAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse 将 AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse(v *AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse.Put(v)
}
