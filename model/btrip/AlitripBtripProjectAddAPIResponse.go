package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectAddAPIResponse 添加项目 API返回值
// alitrip.btrip.project.add
//
// 添加项目
type AlitripBtripProjectAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripProjectAddAPIResponseModel
}

// AlitripBtripProjectAddAPIResponseModel is 添加项目 成功返回结果
type AlitripBtripProjectAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_project_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
