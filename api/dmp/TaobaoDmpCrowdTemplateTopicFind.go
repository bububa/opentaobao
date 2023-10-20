package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// Taobaodmpcrowdtemplatetopicfind 平台精选榜单和模版查询接口
// taobao.dmp.crowd.template.topic.find
//
// 查询平台精选榜单和模版信息
func Taobaodmpcrowdtemplatetopicfind(clt *core.SDKClient, req *dmp.TaobaodmpcrowdtemplatetopicfindAPIRequest, session string) (*dmp.TaobaodmpcrowdtemplatetopicfindAPIResponse, error) {
	var resp dmp.TaobaodmpcrowdtemplatetopicfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
