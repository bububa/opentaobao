package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdruploadretail 快易通多融零售上传接口
// alibaba.alihealth.drug.kyt.druploadretail
//
// 快易通多融零售上传接口
func Alibabaalihealthdrugkytdruploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdruploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdruploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdruploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
