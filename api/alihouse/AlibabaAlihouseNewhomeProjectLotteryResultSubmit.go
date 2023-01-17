package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectLotteryResultSubmit 楼盘摇号结果提交
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
func AlibabaAlihouseNewhomeProjectLotteryResultSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
