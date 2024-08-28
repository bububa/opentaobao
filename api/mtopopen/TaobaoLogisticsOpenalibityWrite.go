package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsOpenalibityWrite 为快递公司提供的物流信息通用写入接口
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
func TaobaoLogisticsOpenalibityWrite(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoLogisticsOpenalibityWriteAPIRequest, resp *mtopopen.TaobaoLogisticsOpenalibityWriteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
