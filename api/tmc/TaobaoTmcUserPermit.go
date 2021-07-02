package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcUserPermit 为已授权的用户开通消息服务
// taobao.tmc.user.permit
//
// 为已授权的用户开通消息服务，授权消息使用。<br/><span style="color:red">注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic</span>
func TaobaoTmcUserPermit(clt *core.SDKClient, req *tmc.TaobaoTmcUserPermitAPIRequest, session string) (*tmc.TaobaoTmcUserPermitAPIResponse, error) {
	var resp tmc.TaobaoTmcUserPermitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
