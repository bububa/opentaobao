package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadzonelist 批量查询可用广告位列表
// taobao.feedflow.item.adzone.list
//
// 批量查询可用广告位列表
func Taobaofeedflowitemadzonelist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadzonelistAPIRequest, session string) (*feedflow.TaobaofeedflowitemadzonelistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadzonelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
