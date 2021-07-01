package drug

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrSpuQueryAPIResponse
获取标品库标品信息 API返回值
alibaba.alihealth.nr.spu.query

提供给ERP使用的，获取健康标品库信息 */
type AlibabaAlihealthNrSpuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthNrSpuQueryAPIResponseModel
}

// AlibabaAlihealthNrSpuQueryAPIResponseModel is 获取标品库标品信息 成功返回结果
type AlibabaAlihealthNrSpuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_spu_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}
