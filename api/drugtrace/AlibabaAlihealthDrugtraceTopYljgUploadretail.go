package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugtraceTopYljgUploadretail
零售单据上传接口
alibaba.alihealth.drugtrace.top.yljg.uploadretail

快易通多融零售上传接口 */
func AlibabaAlihealthDrugtraceTopYljgUploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgUploadretailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgUploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgUploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
