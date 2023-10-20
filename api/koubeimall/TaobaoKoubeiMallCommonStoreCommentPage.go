package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonstorecommentpage 分页查询门店评论详情信息
// taobao.koubei.mall.common.store.comment.page
//
// 查询口碑综合体内的门店评论信息
func Taobaokoubeimallcommonstorecommentpage(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonstorecommentpageAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonstorecommentpageAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonstorecommentpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
