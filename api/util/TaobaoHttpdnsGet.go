package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoHttpdnsGet TOPDNS配置
// taobao.httpdns.get
//
// 获取TOP DNS配置
func TaobaoHttpdnsGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoHttpdnsGetAPIRequest, resp *util.TaobaoHttpdnsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
