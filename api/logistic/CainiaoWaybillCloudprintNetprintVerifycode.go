package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Cainiaowaybillcloudprintnetprintverifycode 打印验证码
// cainiao.waybill.cloudprint.netprint.verifycode
//
// 打印获取验证码
func Cainiaowaybillcloudprintnetprintverifycode(clt *core.SDKClient, req *logistic.CainiaowaybillcloudprintnetprintverifycodeAPIRequest, session string) (*logistic.CainiaowaybillcloudprintnetprintverifycodeAPIResponse, error) {
	var resp logistic.CainiaowaybillcloudprintnetprintverifycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
