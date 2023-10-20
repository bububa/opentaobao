package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqyuploadrelation 关联关系上传
// alibaba.alihealth.drug.kyt.scqy.uploadrelation
//
// 关联关系上传
func Alibabaalihealthdrugkytscqyuploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqyuploadrelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqyuploadrelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqyuploadrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
