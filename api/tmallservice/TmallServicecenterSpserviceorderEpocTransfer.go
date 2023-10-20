package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterspserviceorderepoctransfer 电子保单受保人转移
// tmall.servicecenter.spserviceorder.epoc.transfer
//
// 电子保单受保人转移
func Tmallservicecenterspserviceorderepoctransfer(clt *core.SDKClient, req *tmallservice.TmallservicecenterspserviceorderepoctransferAPIRequest, session string) (*tmallservice.TmallservicecenterspserviceorderepoctransferAPIResponse, error) {
	var resp tmallservice.TmallservicecenterspserviceorderepoctransferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
