package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCollectAPIResponse 工单揽件 API返回值
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
type TmallServicecenterWorkcardCollectAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardCollectAPIResponseModel
}

// TmallServicecenterWorkcardCollectAPIResponseModel is 工单揽件 成功返回结果
type TmallServicecenterWorkcardCollectAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_collect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
