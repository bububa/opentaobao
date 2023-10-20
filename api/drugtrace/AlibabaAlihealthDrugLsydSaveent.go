package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdruglsydsaveent 零售药店往来单位新增
// alibaba.alihealth.drug.lsyd.saveent
//
// 新增往来单位企业记录
func Alibabaalihealthdruglsydsaveent(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdruglsydsaveentAPIRequest, session string) (*drugtrace.AlibabaalihealthdruglsydsaveentAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdruglsydsaveentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
