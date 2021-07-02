package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdzoneRptdailylist 资源包分日数据查询
// taobao.feedflow.item.adzone.rptdailylist
//
// 资源包分日数据查询
func TaobaoFeedflowItemAdzoneRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdzoneRptdailylistAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdzoneRptdailylistAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemAdzoneRptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
