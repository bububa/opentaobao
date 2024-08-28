package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportParterSynccard 阿里体育-卡信息同步接口
// alibaba.alisports.passport.parter.synccard
//
// 运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
func AlibabaAlisportsPassportParterSynccard(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportParterSynccardAPIRequest, resp *alisports.AlibabaAlisportsPassportParterSynccardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
