package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderQueryAPIResponse 服务单列表查询 API返回值
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
type TmallServicecenterSpserviceorderQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderQueryAPIResponseModel
}

// TmallServicecenterSpserviceorderQueryAPIResponseModel is 服务单列表查询 成功返回结果
type TmallServicecenterSpserviceorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
