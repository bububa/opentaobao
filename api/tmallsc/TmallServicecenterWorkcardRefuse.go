package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterworkcardrefuse 买家拒收
// tmall.servicecenter.workcard.refuse
//
// 买家拒收通知接口
func Tmallservicecenterworkcardrefuse(clt *core.SDKClient, req *tmallsc.TmallservicecenterworkcardrefuseAPIRequest, session string) (*tmallsc.TmallservicecenterworkcardrefuseAPIResponse, error) {
	var resp tmallsc.TmallservicecenterworkcardrefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
