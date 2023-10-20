package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleuserpermit 用户appkey授权
// alibaba.idle.user.permit
//
// 闲鱼卖家与服务商关系绑定，用于业务数据分发/授权校验/消息通知鉴权
func Alibabaidleuserpermit(clt *core.SDKClient, req *idle.AlibabaidleuserpermitAPIRequest, session string) (*idle.AlibabaidleuserpermitAPIResponse, error) {
	var resp idle.AlibabaidleuserpermitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
