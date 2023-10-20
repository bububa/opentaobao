package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardRepairprogressUpdate 更新维修进度
// tmall.servicecenter.workcard.repairprogress.update
//
// 提供给外部合作服务商的维修进度更改接口
func TmallServicecenterWorkcardRepairprogressUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardRepairprogressUpdateAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardRepairprogressUpdateAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardRepairprogressUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
