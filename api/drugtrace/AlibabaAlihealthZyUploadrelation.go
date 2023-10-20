package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthzyuploadrelation 中药片关联关系上传
// alibaba.alihealth.zy.uploadrelation
//
// 中药片关联关系上传
func Alibabaalihealthzyuploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthzyuploadrelationAPIRequest, session string) (*drugtrace.AlibabaalihealthzyuploadrelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthzyuploadrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
