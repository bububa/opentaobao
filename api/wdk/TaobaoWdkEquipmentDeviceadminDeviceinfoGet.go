package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentdeviceadmindeviceinfoget 获取五道口设备管理信息
// taobao.wdk.equipment.deviceadmin.deviceinfo.get
//
// 通过仓编码获取五道口设备管理信息
func Taobaowdkequipmentdeviceadmindeviceinfoget(clt *core.SDKClient, req *wdk.TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest, session string) (*wdk.TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
