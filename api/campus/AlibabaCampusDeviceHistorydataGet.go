package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdevicehistorydataget 设备历史数据批量获取
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
func Alibabacampusdevicehistorydataget(clt *core.SDKClient, req *campus.AlibabacampusdevicehistorydatagetAPIRequest, session string) (*campus.AlibabacampusdevicehistorydatagetAPIResponse, error) {
	var resp campus.AlibabacampusdevicehistorydatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
