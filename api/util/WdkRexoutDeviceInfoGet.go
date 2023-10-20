package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Wdkrexoutdeviceinfoget 获取设备详情-外部对接
// wdk.rexout.device.info.get
//
// 获取设备详情-外部对接
func Wdkrexoutdeviceinfoget(clt *core.SDKClient, req *util.WdkrexoutdeviceinfogetAPIRequest, session string) (*util.WdkrexoutdeviceinfogetAPIResponse, error) {
	var resp util.WdkrexoutdeviceinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
