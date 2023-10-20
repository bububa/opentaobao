package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenterworkcardcreate 服务平台工单创建接口
// alibaba.servicecenter.workcard.create
//
// 创建服务平台工单
func Alibabaservicecenterworkcardcreate(clt *core.SDKClient, req *tmallservice.AlibabaservicecenterworkcardcreateAPIRequest, session string) (*tmallservice.AlibabaservicecenterworkcardcreateAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenterworkcardcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
