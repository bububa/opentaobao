package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoNbaddAppointdeliverGetconsigninfo 获取支持定时派送服务发货信息
// cainiao.nbadd.appointdeliver.getconsigninfo
//
// 获取支持定时派送服务发货信息
func CainiaoNbaddAppointdeliverGetconsigninfo(clt *core.SDKClient, req *cainiaolocker.CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest, session string) (*cainiaolocker.CainiaoNbaddAppointdeliverGetconsigninfoAPIResponse, error) {
	var resp cainiaolocker.CainiaoNbaddAppointdeliverGetconsigninfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
