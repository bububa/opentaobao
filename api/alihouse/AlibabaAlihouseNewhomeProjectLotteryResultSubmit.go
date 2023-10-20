package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectlotteryresultsubmit 楼盘摇号结果提交
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
func Alibabaalihousenewhomeprojectlotteryresultsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
