package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemalgocrowdsuggest 单品人群建议出价
// taobao.feedflow.item.algo.crowd.suggest
//
// 给超级推荐的广告主查看建议出价
func Taobaofeedflowitemalgocrowdsuggest(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemalgocrowdsuggestAPIRequest, session string) (*feedflow.TaobaofeedflowitemalgocrowdsuggestAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemalgocrowdsuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
