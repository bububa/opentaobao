package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgeniethirdtelecomautoauth 电信iot自动授权
// alibaba.ailabs.tmallgenie.third.telecom.autoauth
//
// 电信iot自动授权
func Alibabaailabstmallgeniethirdtelecomautoauth(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgeniethirdtelecomautoauthAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgeniethirdtelecomautoauthAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgeniethirdtelecomautoauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
