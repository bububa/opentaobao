package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomercchangestatus 图文草稿状态更新
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
func Alibabaalihousenewhomercchangestatus(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomercchangestatusAPIRequest, session string) (*alihouse.AlibabaalihousenewhomercchangestatusAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomercchangestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
