package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrlistparts 多融查询一个企业的往来单位
// alibaba.alihealth.drug.kyt.dr.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugkytdrlistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrlistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrlistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrlistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
