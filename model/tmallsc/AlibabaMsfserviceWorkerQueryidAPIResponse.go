package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamsfserviceworkerqueryidAPIResponse 查询师傅workerid API返回值
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
type AlibabamsfserviceworkerqueryidAPIResponse struct {
	model.CommonResponse
	AlibabamsfserviceworkerqueryidAPIResponseModel
}

// AlibabamsfserviceworkerqueryidAPIResponseModel is 查询师傅workerid 成功返回结果
type AlibabamsfserviceworkerqueryidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_msfservice_worker_queryid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabamsfserviceworkerqueryidResult `json:"result,omitempty" xml:"result,omitempty"`
}
