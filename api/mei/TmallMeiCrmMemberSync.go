package mei

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallMeiCrmMemberSync 同步推送会员信息
// tmall.mei.crm.member.sync
//
// 品牌方通过该api主动推送会员信息。使用场景包括
// 1.用户在线上注册后，线下补充信息后，同步到线上。
// 2.其他情况的主动推送变更。
func TmallMeiCrmMemberSync(ctx context.Context, clt *core.SDKClient, req *mei.TmallMeiCrmMemberSyncAPIRequest, resp *mei.TmallMeiCrmMemberSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
