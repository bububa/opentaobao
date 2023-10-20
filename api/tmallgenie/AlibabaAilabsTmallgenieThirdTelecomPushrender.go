package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgeniethirdtelecompushrender 电信-推送拉起设备应用
// alibaba.ailabs.tmallgenie.third.telecom.pushrender
//
// 电信-推送拉起设备应用
func Alibabaailabstmallgeniethirdtelecompushrender(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgeniethirdtelecompushrenderAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgeniethirdtelecompushrenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
