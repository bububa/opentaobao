package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytrecordinfo 快易通健康检查
// alibaba.alihealth.drug.kyt.recordinfo
//
// 快易通健康检查
func Alibabaalihealthdrugkytrecordinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytrecordinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytrecordinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytrecordinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
