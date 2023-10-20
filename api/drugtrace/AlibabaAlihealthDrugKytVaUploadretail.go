package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytvauploadretail 零售单据上传接口
// alibaba.alihealth.drug.kyt.va.uploadretail
//
// 零售上传单据信息接口
func Alibabaalihealthdrugkytvauploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytvauploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytvauploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytvauploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
