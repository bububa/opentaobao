package tmalltrend

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallTrendStyleBasicinfoUpload 3D款式基本信息同步API
// tmall.trend.style.basicinfo.upload
//
// 3D款式基本信息同步至天猫趋势中心
func TmallTrendStyleBasicinfoUpload(ctx context.Context, clt *core.SDKClient, req *tmalltrend.TmallTrendStyleBasicinfoUploadAPIRequest, resp *tmalltrend.TmallTrendStyleBasicinfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
