package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeList 授权用户的公益时记录查询
// alibaba.charity.charitytime.list
//
// 查询授权用户的公益时记录
func AlibabaCharityCharitytimeList(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeListAPIRequest, resp *charity.AlibabaCharityCharitytimeListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
