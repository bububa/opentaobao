package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenteridentifytaskcreate 服务商创建核销单
// tmall.servicecenter.identifytask.create
//
// 服务商调用该接口进行创建核销单操作
func Tmallservicecenteridentifytaskcreate(clt *core.SDKClient, req *tmallservice.TmallservicecenteridentifytaskcreateAPIRequest, session string) (*tmallservice.TmallservicecenteridentifytaskcreateAPIResponse, error) {
	var resp tmallservice.TmallservicecenteridentifytaskcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
