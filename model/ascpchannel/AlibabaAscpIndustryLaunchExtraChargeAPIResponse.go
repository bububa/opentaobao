package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryLaunchExtraChargeAPIResponse 阿里巴巴.行业.增加费用.服务商发起 API返回值
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
type AlibabaAscpIndustryLaunchExtraChargeAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryLaunchExtraChargeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryLaunchExtraChargeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryLaunchExtraChargeAPIResponseModel).Reset()
}

// AlibabaAscpIndustryLaunchExtraChargeAPIResponseModel is 阿里巴巴.行业.增加费用.服务商发起 成功返回结果
type AlibabaAscpIndustryLaunchExtraChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_launch_extra_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryLaunchExtraChargeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpIndustryLaunchExtraChargeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryLaunchExtraChargeAPIResponse)
	},
}

// GetAlibabaAscpIndustryLaunchExtraChargeAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryLaunchExtraChargeAPIResponse
func GetAlibabaAscpIndustryLaunchExtraChargeAPIResponse() *AlibabaAscpIndustryLaunchExtraChargeAPIResponse {
	return poolAlibabaAscpIndustryLaunchExtraChargeAPIResponse.Get().(*AlibabaAscpIndustryLaunchExtraChargeAPIResponse)
}

// ReleaseAlibabaAscpIndustryLaunchExtraChargeAPIResponse 将 AlibabaAscpIndustryLaunchExtraChargeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryLaunchExtraChargeAPIResponse(v *AlibabaAscpIndustryLaunchExtraChargeAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryLaunchExtraChargeAPIResponse.Put(v)
}
