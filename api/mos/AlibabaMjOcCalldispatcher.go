package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcCalldispatcher 呼叫运力
// alibaba.mj.oc.calldispatcher
//
// 定时达呼叫运力接口
func AlibabaMjOcCalldispatcher(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcCalldispatcherAPIRequest, resp *mos.AlibabaMjOcCalldispatcherAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
