package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcuserpermit 为已授权的用户开通消息服务
// taobao.tmc.user.permit
//
// 为已授权的用户开通消息服务，授权消息使用。&lt;br/&gt;&lt;span style=&#34;color:red&#34;&gt;注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic&lt;/span&gt;
func Taobaotmcuserpermit(clt *core.SDKClient, req *tmc.TaobaotmcuserpermitAPIRequest, session string) (*tmc.TaobaotmcuserpermitAPIResponse, error) {
	var resp tmc.TaobaotmcuserpermitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
