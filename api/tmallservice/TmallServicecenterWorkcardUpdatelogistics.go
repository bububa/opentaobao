package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardupdatelogistics 更新物流进度
// tmall.servicecenter.workcard.updatelogistics
//
// 提供给外部合作服务商的物流进度更改接口
func Tmallservicecenterworkcardupdatelogistics(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardupdatelogisticsAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardupdatelogisticsAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardupdatelogisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
