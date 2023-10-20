package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectLotteryResultSubmit 楼盘摇号结果提交
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
func AlibabaAlihouseNewhomeProjectLotteryResultSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
