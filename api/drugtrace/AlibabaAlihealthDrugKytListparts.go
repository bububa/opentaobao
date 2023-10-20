package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytlistparts 查询往来单位列表
// alibaba.alihealth.drug.kyt.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugkytlistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytlistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytlistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytlistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
