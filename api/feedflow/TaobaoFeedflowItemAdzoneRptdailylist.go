package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadzonerptdailylist 资源包分日数据查询
// taobao.feedflow.item.adzone.rptdailylist
//
// 资源包分日数据查询
func Taobaofeedflowitemadzonerptdailylist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadzonerptdailylistAPIRequest, session string) (*feedflow.TaobaofeedflowitemadzonerptdailylistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadzonerptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
