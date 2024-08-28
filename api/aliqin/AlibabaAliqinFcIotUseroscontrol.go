package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotUseroscontrol 物联卡用户管理停开机功能
// alibaba.aliqin.fc.iot.useroscontrol
//
// 物联网针对用户级管理停开支持
func AlibabaAliqinFcIotUseroscontrol(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotUseroscontrolAPIRequest, resp *aliqin.AlibabaAliqinFcIotUseroscontrolAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
