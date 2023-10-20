package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmctopicgroupdelete 删除消息topic分组路由
// taobao.tmc.topic.group.delete
//
// 删除根据topic名称路由消息到不同的分组关系
func Taobaotmctopicgroupdelete(clt *core.SDKClient, req *tmc.TaobaotmctopicgroupdeleteAPIRequest, session string) (*tmc.TaobaotmctopicgroupdeleteAPIResponse, error) {
	var resp tmc.TaobaotmctopicgroupdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
