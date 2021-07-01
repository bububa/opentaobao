package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletePaperformatAPIResponse
大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat API返回值
alibaba.damai.mev.open.delete.paperformat

deletePaperFormat */
type AlibabaDamaiMevOpenDeletePaperformatAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel
}

// AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel is 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat 成功返回结果
type AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_delete_paperformat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletePaperformatResult `json:"result,omitempty" xml:"result,omitempty"`
}
