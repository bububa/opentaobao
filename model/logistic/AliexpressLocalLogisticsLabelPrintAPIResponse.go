package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticslabelprintAPIResponse print label API返回值
// aliexpress.local.logistics.label.print
//
// print label
type AliexpresslocallogisticslabelprintAPIResponse struct {
	model.CommonResponse
	AliexpresslocallogisticslabelprintAPIResponseModel
}

// AliexpresslocallogisticslabelprintAPIResponseModel is print label 成功返回结果
type AliexpresslocallogisticslabelprintAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistics_label_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error message
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// error code
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// data
	Data *LabelDto `json:"data,omitempty" xml:"data,omitempty"`
	// is success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
