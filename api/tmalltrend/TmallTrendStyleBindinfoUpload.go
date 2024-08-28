package tmalltrend

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallTrendStyleBindinfoUpload 趋势词&款式绑定信息同步API
// tmall.trend.style.bindinfo.upload
//
// 趋势词&amp;款式(服饰行业)绑定信息同步至平台
func TmallTrendStyleBindinfoUpload(ctx context.Context, clt *core.SDKClient, req *tmalltrend.TmallTrendStyleBindinfoUploadAPIRequest, resp *tmalltrend.TmallTrendStyleBindinfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
