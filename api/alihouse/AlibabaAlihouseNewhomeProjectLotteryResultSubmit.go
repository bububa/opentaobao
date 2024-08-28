package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectLotteryResultSubmit 楼盘摇号结果提交
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
func AlibabaAlihouseNewhomeProjectLotteryResultSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
