package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsiotbusinessrecipestepinsertorupdate 插入或更新食谱步骤
// alibaba.ailabs.iot.business.recipestep.insertorupdate
//
// 插入或更新食谱步骤
func Alibabaailabsiotbusinessrecipestepinsertorupdate(clt *core.SDKClient, req *iot.AlibabaailabsiotbusinessrecipestepinsertorupdateAPIRequest, session string) (*iot.AlibabaailabsiotbusinessrecipestepinsertorupdateAPIResponse, error) {
	var resp iot.AlibabaailabsiotbusinessrecipestepinsertorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
