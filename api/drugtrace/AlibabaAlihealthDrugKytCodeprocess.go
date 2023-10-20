package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytcodeprocess 关联关系处理查询
// alibaba.alihealth.drug.kyt.codeprocess
//
// 关联关系处理查询
func Alibabaalihealthdrugkytcodeprocess(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytcodeprocessAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytcodeprocessAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytcodeprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
