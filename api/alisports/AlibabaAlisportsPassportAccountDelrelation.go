package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

/* AlibabaAlisportsPassportAccountDelrelation
阿里体育会员系统--取消三方关联接口
alibaba.alisports.passport.account.delrelation

阿里体育会员系统--取消三方关联接口 */
func AlibabaAlisportsPassportAccountDelrelation(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountDelrelationAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAccountDelrelationAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportAccountDelrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
