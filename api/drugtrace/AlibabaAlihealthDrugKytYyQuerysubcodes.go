package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytyyquerysubcodes 查询一个码的所有子码
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
func Alibabaalihealthdrugkytyyquerysubcodes(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytyyquerysubcodesAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytyyquerysubcodesAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytyyquerysubcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
