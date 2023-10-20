package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentDeviceadminDeviceinfoGet 获取五道口设备管理信息
// taobao.wdk.equipment.deviceadmin.deviceinfo.get
//
// 通过仓编码获取五道口设备管理信息
func TaobaoWdkEquipmentDeviceadminDeviceinfoGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest, resp *wdk.TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
