package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmsdclaimqueryAPIResponse 查询待理赔工单数据接口 API返回值
// tmall.msd.claim.query
//
// 查询待理赔工单数据接口
type TmallmsdclaimqueryAPIResponse struct {
	model.CommonResponse
	TmallmsdclaimqueryAPIResponseModel
}

// TmallmsdclaimqueryAPIResponseModel is 查询待理赔工单数据接口 成功返回结果
type TmallmsdclaimqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msd_claim_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
