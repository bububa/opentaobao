package idleparttime

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleparttimejobsyncAPIResponse 兼职岗位同步 API返回值
// alibaba.idle.parttime.jobsync
//
// 服务商同步岗位信息给闲鱼
type AlibabaidleparttimejobsyncAPIResponse struct {
	model.CommonResponse
	AlibabaidleparttimejobsyncAPIResponseModel
}

// AlibabaidleparttimejobsyncAPIResponseModel is 兼职岗位同步 成功返回结果
type AlibabaidleparttimejobsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_parttime_jobsync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
