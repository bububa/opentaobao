package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterWorkcardServiceprogressUpdate 更新服务进度
// alibaba.servicecenter.workcard.serviceprogress.update
//
// 提供给外部合作服务商更新服务进度的接口
func AlibabaServicecenterWorkcardServiceprogressUpdate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterWorkcardServiceprogressUpdateAPIRequest, session string) (*tmallservice.AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse, error) {
	var resp tmallservice.AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
