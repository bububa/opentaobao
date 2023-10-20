package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytquerycodeactive 查询码是否激活
// alibaba.alihealth.drug.kyt.querycodeactive
//
// 查询码是否激活
func Alibabaalihealthdrugkytquerycodeactive(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytquerycodeactiveAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytquerycodeactiveAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytquerycodeactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
