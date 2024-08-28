package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportAsyncCreateDownloadTask 创建异步下载任务
// taobao.universalbp.report.async.create.download.task
//
// 入参报表查询信息，出参下载任务id
func TaobaoUniversalbpReportAsyncCreateDownloadTask(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest, resp *simba.TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
