package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesSort 系列品-商品排序
// alibaba.wdk.series.sort
//
// 系列品商品变更-商品排序
func AlibabaWdkSeriesSort(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSeriesSortAPIRequest, resp *wdk.AlibabaWdkSeriesSortAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
