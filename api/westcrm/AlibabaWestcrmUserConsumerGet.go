package westcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmUserConsumerGet 获取指定用户的消费总额
// alibaba.westcrm.user.consumer.get
//
// 获取指定用户的消费总额
func AlibabaWestcrmUserConsumerGet(ctx context.Context, clt *core.SDKClient, req *westcrm.AlibabaWestcrmUserConsumerGetAPIRequest, resp *westcrm.AlibabaWestcrmUserConsumerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
