package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthtracecodesellerbillrootcodeget 获取最外层包装码
// alibaba.alihealth.tracecodeseller.bill.rootcode.get
//
// 获取最外层包装码
func Alibabaalihealthtracecodesellerbillrootcodeget(clt *core.SDKClient, req *drugtrace.AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest, session string) (*drugtrace.AlibabaalihealthtracecodesellerbillrootcodegetAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthtracecodesellerbillrootcodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
