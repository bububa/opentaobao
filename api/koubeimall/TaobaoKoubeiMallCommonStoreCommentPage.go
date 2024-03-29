package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStoreCommentPage 分页查询门店评论详情信息
// taobao.koubei.mall.common.store.comment.page
//
// 查询口碑综合体内的门店评论信息
func TaobaoKoubeiMallCommonStoreCommentPage(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreCommentPageAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonStoreCommentPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
