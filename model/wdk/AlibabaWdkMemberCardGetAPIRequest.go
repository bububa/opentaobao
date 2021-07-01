package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMemberCardGetAPIRequest
查询会员卡信息 API请求
alibaba.wdk.member.card.get

根据会员卡查询会员信息 */
type AlibabaWdkMemberCardGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_memberQuery *MemberQueryRequest
}

// New
