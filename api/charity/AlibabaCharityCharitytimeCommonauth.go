package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeCommonauth 通用授权
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
func AlibabaCharityCharitytimeCommonauth(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeCommonauthAPIRequest, session string) (*charity.AlibabaCharityCharitytimeCommonauthAPIResponse, error) {
	var resp charity.AlibabaCharityCharitytimeCommonauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
