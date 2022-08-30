package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuSearchAPIResponse 供应链中台货品搜索接口 API返回值
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
type AlibabaAscpCnskuSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAscpCnskuSearchAPIResponseModel
}

// AlibabaAscpCnskuSearchAPIResponseModel is 供应链中台货品搜索接口 成功返回结果
type AlibabaAscpCnskuSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_cnsku_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CnskuResult `json:"result,omitempty" xml:"result,omitempty"`
}
