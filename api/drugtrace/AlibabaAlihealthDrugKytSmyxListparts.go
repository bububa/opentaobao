package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsmyxlistparts 药店查询往来单位
// alibaba.alihealth.drug.kyt.smyx.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugkytsmyxlistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsmyxlistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsmyxlistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsmyxlistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
