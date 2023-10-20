package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytyylistparts 查询往来单位
// alibaba.alihealth.drug.kyt.yy.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugkytyylistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytyylistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytyylistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytyylistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
