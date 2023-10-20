package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimecommonauth 通用授权
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
func Alibabacharitycharitytimecommonauth(clt *core.SDKClient, req *charity.AlibabacharitycharitytimecommonauthAPIRequest, session string) (*charity.AlibabacharitycharitytimecommonauthAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimecommonauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
