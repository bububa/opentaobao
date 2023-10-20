package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytuploadrelation 关联关系上传
// alibaba.alihealth.drug.kyt.uploadrelation
//
// 关联关系上传
func Alibabaalihealthdrugkytuploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytuploadrelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytuploadrelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytuploadrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
