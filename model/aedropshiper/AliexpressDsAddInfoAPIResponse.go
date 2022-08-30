package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsAddInfoAPIResponse 上报DS信息 API返回值
// aliexpress.ds.add.info
//
// ISV用户上报下游DS信息
type AliexpressDsAddInfoAPIResponse struct {
	model.CommonResponse
	AliexpressDsAddInfoAPIResponseModel
}

// AliexpressDsAddInfoAPIResponseModel is 上报DS信息 成功返回结果
type AliexpressDsAddInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_add_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Error message
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// Error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// Result object
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
