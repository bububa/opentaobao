package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmembercardget 查询会员卡信息
// alibaba.wdk.member.card.get
//
// 根据会员卡查询会员信息
func Alibabawdkmembercardget(clt *core.SDKClient, req *wdk.AlibabawdkmembercardgetAPIRequest, session string) (*wdk.AlibabawdkmembercardgetAPIResponse, error) {
	var resp wdk.AlibabawdkmembercardgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
