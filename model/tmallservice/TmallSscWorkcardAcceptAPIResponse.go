package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallsscworkcardacceptAPIResponse 服务商接单完成 API返回值
// tmall.ssc.workcard.accept
//
// 工单完结
type TmallsscworkcardacceptAPIResponse struct {
	model.CommonResponse
	TmallsscworkcardacceptAPIResponseModel
}

// TmallsscworkcardacceptAPIResponseModel is 服务商接单完成 成功返回结果
type TmallsscworkcardacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ssc_workcard_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接单结果
	Result *TmallsscworkcardacceptResult `json:"result,omitempty" xml:"result,omitempty"`
}
