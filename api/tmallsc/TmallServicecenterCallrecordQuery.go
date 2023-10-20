package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecentercallrecordquery 天猫服务平台服务商查询通话记录接口
// tmall.servicecenter.callrecord.query
//
// 天猫服务平台服务商查询通话记录接口
func Tmallservicecentercallrecordquery(clt *core.SDKClient, req *tmallsc.TmallservicecentercallrecordqueryAPIRequest, session string) (*tmallsc.TmallservicecentercallrecordqueryAPIResponse, error) {
	var resp tmallsc.TmallservicecentercallrecordqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
