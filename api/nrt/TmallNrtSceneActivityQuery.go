package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtsceneactivityquery 喵零场景活动查询
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
func Tmallnrtsceneactivityquery(clt *core.SDKClient, req *nrt.TmallnrtsceneactivityqueryAPIRequest, session string) (*nrt.TmallnrtsceneactivityqueryAPIResponse, error) {
	var resp nrt.TmallnrtsceneactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
