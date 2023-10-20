package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytweslistparts 查询往来单位列表
// alibaba.alihealth.drug.kyt.wes.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugkytweslistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytweslistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytweslistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytweslistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
