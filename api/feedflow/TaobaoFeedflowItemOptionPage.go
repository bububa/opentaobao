package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemoptionpage 分页查询定向标签列表
// taobao.feedflow.item.option.page
//
// 分页查询定向标签列表
func Taobaofeedflowitemoptionpage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemoptionpageAPIRequest, session string) (*feedflow.TaobaofeedflowitemoptionpageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemoptionpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
