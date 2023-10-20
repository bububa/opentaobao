package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Alipayxiaodaiuserpermit 阿里金融用户授权
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
func Alipayxiaodaiuserpermit(clt *core.SDKClient, req *tmc.AlipayxiaodaiuserpermitAPIRequest, session string) (*tmc.AlipayxiaodaiuserpermitAPIResponse, error) {
	var resp tmc.AlipayxiaodaiuserpermitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
