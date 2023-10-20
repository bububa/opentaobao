package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMsfserviceWorkerQueryidAPIResponse 查询师傅workerid API返回值
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
type AlibabaMsfserviceWorkerQueryidAPIResponse struct {
	model.CommonResponse
	AlibabaMsfserviceWorkerQueryidAPIResponseModel
}

// AlibabaMsfserviceWorkerQueryidAPIResponseModel is 查询师傅workerid 成功返回结果
type AlibabaMsfserviceWorkerQueryidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_msfservice_worker_queryid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabaMsfserviceWorkerQueryidResult `json:"result,omitempty" xml:"result,omitempty"`
}
