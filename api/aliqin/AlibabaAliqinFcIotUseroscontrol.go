package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotuseroscontrol 物联卡用户管理停开机功能
// alibaba.aliqin.fc.iot.useroscontrol
//
// 物联网针对用户级管理停开支持
func Alibabaaliqinfciotuseroscontrol(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotuseroscontrolAPIRequest, session string) (*aliqin.AlibabaaliqinfciotuseroscontrolAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotuseroscontrolAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
