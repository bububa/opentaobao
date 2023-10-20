package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportsdatacenterdatasynctreadmill 阿里体育同步跑步机设备数据
// alibaba.alisports.datacenter.datasync.treadmill
//
// 合作方向阿里体育同步跑步机设备的数据
func Alibabaalisportsdatacenterdatasynctreadmill(clt *core.SDKClient, req *alisports.AlibabaalisportsdatacenterdatasynctreadmillAPIRequest, session string) (*alisports.AlibabaalisportsdatacenterdatasynctreadmillAPIResponse, error) {
	var resp alisports.AlibabaalisportsdatacenterdatasynctreadmillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
