package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgContentGetAPIResponse
百川内容平台内容获取 API返回值
alibaba.baichuan.ctg.content.get

百川内容平台内容获取 */
type AlibabaBaichuanCtgContentGetAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanCtgContentGetAPIResponseModel
}

// AlibabaBaichuanCtgContentGetAPIResponseModel is 百川内容平台内容获取 成功返回结果
type AlibabaBaichuanCtgContentGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_content_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// hasNext
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// errorCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// data
	DataList []AlibabaBaichuanCtgContentGetData `json:"data_list,omitempty" xml:"data_list>alibaba_baichuan_ctg_content_get_data,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
