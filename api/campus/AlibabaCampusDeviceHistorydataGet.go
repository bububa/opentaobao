package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceHistorydataGet 设备历史数据批量获取
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
func AlibabaCampusDeviceHistorydataGet(clt *core.SDKClient, req *campus.AlibabaCampusDeviceHistorydataGetAPIRequest, resp *campus.AlibabaCampusDeviceHistorydataGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
