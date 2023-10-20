package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydUploadretail 零售单据上传接口
// alibaba.alihealth.drugtrace.top.lsyd.uploadretail
//
// 快易通多融零售上传接口
func AlibabaAlihealthDrugtraceTopLsydUploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
