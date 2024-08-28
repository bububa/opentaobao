package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaCfdaXtptAppGetshowurl 协同平台码查询页面url
// alibaba.cfda.xtpt.app.getshowurl
//
// 协同平台码查询页面url
func AlibabaCfdaXtptAppGetshowurl(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaCfdaXtptAppGetshowurlAPIRequest, resp *drugtrace.AlibabaCfdaXtptAppGetshowurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
