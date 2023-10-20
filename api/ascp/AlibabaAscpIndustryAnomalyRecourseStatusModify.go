package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabaascpindustryanomalyrecoursestatusmodify 送货入户并安装投诉工单状态变更
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
func Alibabaascpindustryanomalyrecoursestatusmodify(clt *core.SDKClient, req *ascp.AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest, session string) (*ascp.AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponse, error) {
	var resp ascp.AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
