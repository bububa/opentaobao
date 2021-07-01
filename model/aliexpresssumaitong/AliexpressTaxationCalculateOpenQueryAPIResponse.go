package aliexpresssumaitong

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTaxationCalculateOpenQueryAPIResponse
关务所需的申报清关字段 API返回值
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段 */
type AliexpressTaxationCalculateOpenQueryAPIResponse struct {
	model.CommonResponse
	AliexpressTaxationCalculateOpenQueryAPIResponseModel
}

// AliexpressTaxationCalculateOpenQueryAPIResponseModel is 关务所需的申报清关字段 成功返回结果
type AliexpressTaxationCalculateOpenQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_taxation_calculate_open_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}
