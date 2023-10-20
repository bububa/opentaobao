package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgrouppage 查询单元列表
// taobao.feedflow.item.adgroup.page
//
// 通过计划id查询单元信息
func Taobaofeedflowitemadgrouppage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgrouppageAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgrouppageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgrouppageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
