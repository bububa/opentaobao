package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterspserviceorderquery 服务单列表查询
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
func Tmallservicecenterspserviceorderquery(clt *core.SDKClient, req *tmallservice.TmallservicecenterspserviceorderqueryAPIRequest, session string) (*tmallservice.TmallservicecenterspserviceorderqueryAPIResponse, error) {
	var resp tmallservice.TmallservicecenterspserviceorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
