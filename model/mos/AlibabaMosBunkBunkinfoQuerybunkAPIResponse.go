package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosbunkbunkinfoquerybunkAPIResponse 根据合同号查询铺位信息 API返回值
// alibaba.mos.bunk.bunkinfo.querybunk
//
// 根据合同号查询铺位信息
type AlibabamosbunkbunkinfoquerybunkAPIResponse struct {
	model.CommonResponse
	AlibabamosbunkbunkinfoquerybunkAPIResponseModel
}

// AlibabamosbunkbunkinfoquerybunkAPIResponseModel is 根据合同号查询铺位信息 成功返回结果
type AlibabamosbunkbunkinfoquerybunkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_bunk_bunkinfo_querybunk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamosbunkbunkinfoquerybunkResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
