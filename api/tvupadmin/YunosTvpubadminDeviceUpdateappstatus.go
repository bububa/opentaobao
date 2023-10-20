package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceUpdateappstatus 更新应用版本审核状态
// yunos.tvpubadmin.device.updateappstatus
//
// 更新应用版本审核状态
func YunosTvpubadminDeviceUpdateappstatus(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceUpdateappstatusAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceUpdateappstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
