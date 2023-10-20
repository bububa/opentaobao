package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthZyUploadrelation 中药片关联关系上传
// alibaba.alihealth.zy.uploadrelation
//
// 中药片关联关系上传
func AlibabaAlihealthZyUploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthZyUploadrelationAPIRequest, resp *drugtrace.AlibabaAlihealthZyUploadrelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
