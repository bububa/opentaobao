package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingUnbundleAPIResponse 商货关联解绑 API返回值
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
type AlibabaDchainAoxiangItemmappingUnbundleAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel
}

// AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel is 商货关联解绑 成功返回结果
type AlibabaDchainAoxiangItemmappingUnbundleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_itemmapping_unbundle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	UnbundleItemmappingResponse *UnbundleItemmappingResponse `json:"unbundle_itemmapping_response,omitempty" xml:"unbundle_itemmapping_response,omitempty"`
}
