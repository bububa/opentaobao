package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenteridentifytaskcreate 创建核销单
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
func Alibabaservicecenteridentifytaskcreate(clt *core.SDKClient, req *tmallservice.AlibabaservicecenteridentifytaskcreateAPIRequest, session string) (*tmallservice.AlibabaservicecenteridentifytaskcreateAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenteridentifytaskcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
