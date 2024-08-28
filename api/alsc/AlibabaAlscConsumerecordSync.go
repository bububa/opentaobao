package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscConsumerecordSync 外域订单同步本地生活消费记录
// alibaba.alsc.consumerecord.sync
//
// 外部第三方将本地生活app端下单数据同步到本地生活消费记录中心
func AlibabaAlscConsumerecordSync(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscConsumerecordSyncAPIRequest, resp *alsc.AlibabaAlscConsumerecordSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
