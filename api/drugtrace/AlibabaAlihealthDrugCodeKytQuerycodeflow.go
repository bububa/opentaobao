package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytquerycodeflow 码流向查询
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
func Alibabaalihealthdrugcodekytquerycodeflow(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytquerycodeflowAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytquerycodeflowAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytquerycodeflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
