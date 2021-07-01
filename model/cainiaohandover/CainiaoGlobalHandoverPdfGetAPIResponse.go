package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverPdfGetAPIResponse
获取面单PDF文件数据 API返回值
cainiao.global.handover.pdf.get

返回指定大包面单的PDF文件数据 */
type CainiaoGlobalHandoverPdfGetAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverPdfGetAPIResponseModel
}

// CainiaoGlobalHandoverPdfGetAPIResponseModel is 获取面单PDF文件数据 成功返回结果
type CainiaoGlobalHandoverPdfGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_pdf_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
