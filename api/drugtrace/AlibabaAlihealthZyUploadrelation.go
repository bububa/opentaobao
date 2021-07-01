package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthZyUploadrelation
中药片关联关系上传
alibaba.alihealth.zy.uploadrelation

中药片关联关系上传 */
func AlibabaAlihealthZyUploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthZyUploadrelationAPIRequest, session string) (*drugtrace.AlibabaAlihealthZyUploadrelationAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthZyUploadrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
