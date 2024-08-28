package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelCityCoordinatesBatchDownload 下载飞猪国际城市结果
// taobao.xhotel.city.coordinates.batch.download
//
// 给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
func TaobaoXhotelCityCoordinatesBatchDownload(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest, resp *xhotelitem.TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
