package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustrylaunchextrachargeAPIResponse 阿里巴巴.行业.增加费用.服务商发起 API返回值
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
type AlibabaascpindustrylaunchextrachargeAPIResponse struct {
	model.CommonResponse
	AlibabaascpindustrylaunchextrachargeAPIResponseModel
}

// AlibabaascpindustrylaunchextrachargeAPIResponseModel is 阿里巴巴.行业.增加费用.服务商发起 成功返回结果
type AlibabaascpindustrylaunchextrachargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_launch_extra_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
