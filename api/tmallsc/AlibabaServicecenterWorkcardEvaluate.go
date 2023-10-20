package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabaservicecenterworkcardevaluate 服务商售后鉴定服务
// alibaba.servicecenter.workcard.evaluate
//
// 服务商售后鉴定服务,提供给服务商针对售后场景上门鉴定服务，鉴定成功则服务商完成履约，鉴定失败则取消工单
func Alibabaservicecenterworkcardevaluate(clt *core.SDKClient, req *tmallsc.AlibabaservicecenterworkcardevaluateAPIRequest, session string) (*tmallsc.AlibabaservicecenterworkcardevaluateAPIResponse, error) {
	var resp tmallsc.AlibabaservicecenterworkcardevaluateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
