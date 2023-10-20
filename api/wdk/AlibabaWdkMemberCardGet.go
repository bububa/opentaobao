package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMemberCardGet 查询会员卡信息
// alibaba.wdk.member.card.get
//
// 根据会员卡查询会员信息
func AlibabaWdkMemberCardGet(clt *core.SDKClient, req *wdk.AlibabaWdkMemberCardGetAPIRequest, resp *wdk.AlibabaWdkMemberCardGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
