package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportaccountdelrelation 阿里体育会员系统--取消三方关联接口
// alibaba.alisports.passport.account.delrelation
//
// 阿里体育会员系统--取消三方关联接口
func Alibabaalisportspassportaccountdelrelation(clt *core.SDKClient, req *alisports.AlibabaalisportspassportaccountdelrelationAPIRequest, session string) (*alisports.AlibabaalisportspassportaccountdelrelationAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportaccountdelrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
