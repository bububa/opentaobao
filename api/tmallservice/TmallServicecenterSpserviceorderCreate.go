package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterspserviceordercreate 服务单创建
// tmall.servicecenter.spserviceorder.create
//
// 服务单创建
func Tmallservicecenterspserviceordercreate(clt *core.SDKClient, req *tmallservice.TmallservicecenterspserviceordercreateAPIRequest, session string) (*tmallservice.TmallservicecenterspserviceordercreateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterspserviceordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
