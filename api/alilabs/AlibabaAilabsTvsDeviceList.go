package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstvsdevicelist 获取TVS设备列表
// alibaba.ailabs.tvs.device.list
//
// 获取用户所绑定的TVS设备列表
func Alibabaailabstvsdevicelist(clt *core.SDKClient, req *alilabs.AlibabaailabstvsdevicelistAPIRequest, session string) (*alilabs.AlibabaailabstvsdevicelistAPIResponse, error) {
	var resp alilabs.AlibabaailabstvsdevicelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
