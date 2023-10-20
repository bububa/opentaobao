package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardreserve 工单预约
// tmall.servicecenter.workcard.reserve
//
// 服务工单更新通用接口
func Tmallservicecenterworkcardreserve(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardreserveAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardreserveAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardreserveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
