package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// AlipayXiaodaiUserPermit 阿里金融用户授权
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
func AlipayXiaodaiUserPermit(ctx context.Context, clt *core.SDKClient, req *tmc.AlipayXiaodaiUserPermitAPIRequest, resp *tmc.AlipayXiaodaiUserPermitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
