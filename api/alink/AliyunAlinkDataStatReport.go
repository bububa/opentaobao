package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AliyunAlinkDataStatReport 外部离线统计数据上报
// aliyun.alink.data.stat.report
//
// 外部合作厂商上报设备的明细数据，或者离线统计数据。
func AliyunAlinkDataStatReport(clt *core.SDKClient, req *alink.AliyunAlinkDataStatReportAPIRequest, session string) (*alink.AliyunAlinkDataStatReportAPIResponse, error) {
	var resp alink.AliyunAlinkDataStatReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
