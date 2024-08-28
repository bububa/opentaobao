package alink

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AliyunAlinkDataStatReport 外部离线统计数据上报
// aliyun.alink.data.stat.report
//
// 外部合作厂商上报设备的明细数据，或者离线统计数据。
func AliyunAlinkDataStatReport(ctx context.Context, clt *core.SDKClient, req *alink.AliyunAlinkDataStatReportAPIRequest, resp *alink.AliyunAlinkDataStatReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
