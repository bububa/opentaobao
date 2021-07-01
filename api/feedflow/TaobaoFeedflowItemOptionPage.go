package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

/* TaobaoFeedflowItemOptionPage
分页查询定向标签列表
taobao.feedflow.item.option.page

分页查询定向标签列表 */
func TaobaoFeedflowItemOptionPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemOptionPageAPIRequest, session string) (*feedflow.TaobaoFeedflowItemOptionPageAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemOptionPageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
