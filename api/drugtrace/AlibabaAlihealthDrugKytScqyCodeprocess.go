package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqycodeprocess 关联关系处理查询
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
func Alibabaalihealthdrugkytscqycodeprocess(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqycodeprocessAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqycodeprocessAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqycodeprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
