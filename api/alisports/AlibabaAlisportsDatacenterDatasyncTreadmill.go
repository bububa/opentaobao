package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDatacenterDatasyncTreadmill 阿里体育同步跑步机设备数据
// alibaba.alisports.datacenter.datasync.treadmill
//
// 合作方向阿里体育同步跑步机设备的数据
func AlibabaAlisportsDatacenterDatasyncTreadmill(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsDatacenterDatasyncTreadmillAPIRequest, resp *alisports.AlibabaAlisportsDatacenterDatasyncTreadmillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
