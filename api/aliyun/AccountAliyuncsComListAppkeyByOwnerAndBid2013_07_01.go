package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

/* AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01
根据渠道id和状态列出appkey
account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01

根据渠道id和状态列出appkey */
func AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest, session string) (*aliyun.AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponse, error) {
	var resp aliyun.AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
