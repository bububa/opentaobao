package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenterworkcardserviceprogressupdate 更新服务进度
// alibaba.servicecenter.workcard.serviceprogress.update
//
// 提供给外部合作服务商更新服务进度的接口
func Alibabaservicecenterworkcardserviceprogressupdate(clt *core.SDKClient, req *tmallservice.AlibabaservicecenterworkcardserviceprogressupdateAPIRequest, session string) (*tmallservice.AlibabaservicecenterworkcardserviceprogressupdateAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenterworkcardserviceprogressupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
