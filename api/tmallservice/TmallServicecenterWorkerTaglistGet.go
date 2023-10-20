package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkertaglistget 获取工人标签
// tmall.servicecenter.worker.taglist.get
//
// 服务商获取对应工人的标签
func Tmallservicecenterworkertaglistget(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkertaglistgetAPIRequest, session string) (*tmallservice.TmallservicecenterworkertaglistgetAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkertaglistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
