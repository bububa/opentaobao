package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// TaobaoDmpCrowdTemplateTopicFind 平台精选榜单和模版查询接口
// taobao.dmp.crowd.template.topic.find
//
// 查询平台精选榜单和模版信息
func TaobaoDmpCrowdTemplateTopicFind(clt *core.SDKClient, req *dmp.TaobaoDmpCrowdTemplateTopicFindAPIRequest, resp *dmp.TaobaoDmpCrowdTemplateTopicFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
