package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotQryPersoninfo 查询物联卡个人实人认证信息
// alibaba.aliqin.fc.iot.qry.personinfo
//
// 查询物联卡个人实人认证信息
func AlibabaAliqinFcIotQryPersoninfo(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotQryPersoninfoAPIRequest, resp *aliqin.AlibabaAliqinFcIotQryPersoninfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
