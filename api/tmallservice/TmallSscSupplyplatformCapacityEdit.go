package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallsscsupplyplatformcapacityedit 容量编辑
// tmall.ssc.supplyplatform.capacity.edit
//
// 容量编辑
func Tmallsscsupplyplatformcapacityedit(clt *core.SDKClient, req *tmallservice.TmallsscsupplyplatformcapacityeditAPIRequest, session string) (*tmallservice.TmallsscsupplyplatformcapacityeditAPIResponse, error) {
	var resp tmallservice.TmallsscsupplyplatformcapacityeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
