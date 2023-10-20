package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanctgcontentgetAPIResponse 百川内容平台内容获取 API返回值
// alibaba.baichuan.ctg.content.get
//
// 百川内容平台内容获取
type AlibababaichuanctgcontentgetAPIResponse struct {
	model.CommonResponse
	AlibababaichuanctgcontentgetAPIResponseModel
}

// AlibababaichuanctgcontentgetAPIResponseModel is 百川内容平台内容获取 成功返回结果
type AlibababaichuanctgcontentgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_content_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	DataList []AlibababaichuanctgcontentgetData `json:"data_list,omitempty" xml:"data_list>alibababaichuanctgcontentget_data,omitempty"`
	// errorMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errorCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// hasNext
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
