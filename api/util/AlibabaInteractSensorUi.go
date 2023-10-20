package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Alibabainteractsensorui 基本ui操作
// alibaba.interact.sensor.ui
//
// Weex 基本UI操作
func Alibabainteractsensorui(clt *core.SDKClient, req *util.AlibabainteractsensoruiAPIRequest, session string) (*util.AlibabainteractsensoruiAPIResponse, error) {
	var resp util.AlibabainteractsensoruiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
