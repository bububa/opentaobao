package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardrepairprogressupdate 更新维修进度
// tmall.servicecenter.workcard.repairprogress.update
//
// 提供给外部合作服务商的维修进度更改接口
func Tmallservicecenterworkcardrepairprogressupdate(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardrepairprogressupdateAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardrepairprogressupdateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardrepairprogressupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
