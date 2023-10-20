package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Alibabainteractlotteryactivityregister 回传抽奖相关参数
// alibaba.interact.lotteryactivity.register
//
// 提供接口供三方应用将数据回传到平台
func Alibabainteractlotteryactivityregister(clt *core.SDKClient, req *mtopopen.AlibabainteractlotteryactivityregisterAPIRequest, session string) (*mtopopen.AlibabainteractlotteryactivityregisterAPIResponse, error) {
	var resp mtopopen.AlibabainteractlotteryactivityregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
