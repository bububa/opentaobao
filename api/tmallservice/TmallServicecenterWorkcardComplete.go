package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardcomplete 工单完结
// tmall.servicecenter.workcard.complete
//
// 工单完结
func Tmallservicecenterworkcardcomplete(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardcompleteAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardcompleteAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardcompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
