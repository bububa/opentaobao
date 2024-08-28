package tmalltrend

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallTrendStyleProduceinfoUpload 款式生产信息同步API
// tmall.trend.style.produceinfo.upload
//
// 款式生产信息同步至平台
func TmallTrendStyleProduceinfoUpload(ctx context.Context, clt *core.SDKClient, req *tmalltrend.TmallTrendStyleProduceinfoUploadAPIRequest, resp *tmalltrend.TmallTrendStyleProduceinfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
