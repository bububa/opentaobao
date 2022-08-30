package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtSceneActivityQuery 喵零场景活动查询
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
func TmallNrtSceneActivityQuery(clt *core.SDKClient, req *nrt.TmallNrtSceneActivityQueryAPIRequest, session string) (*nrt.TmallNrtSceneActivityQueryAPIResponse, error) {
	var resp nrt.TmallNrtSceneActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
