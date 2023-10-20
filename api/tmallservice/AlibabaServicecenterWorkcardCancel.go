package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenterworkcardcancel 服务平台工单取消接口
// alibaba.servicecenter.workcard.cancel
//
// 取消服务工单
func Alibabaservicecenterworkcardcancel(clt *core.SDKClient, req *tmallservice.AlibabaservicecenterworkcardcancelAPIRequest, session string) (*tmallservice.AlibabaservicecenterworkcardcancelAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenterworkcardcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
