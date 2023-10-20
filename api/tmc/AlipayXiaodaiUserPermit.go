package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// AlipayXiaodaiUserPermit 阿里金融用户授权
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
func AlipayXiaodaiUserPermit(clt *core.SDKClient, req *tmc.AlipayXiaodaiUserPermitAPIRequest, resp *tmc.AlipayXiaodaiUserPermitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
