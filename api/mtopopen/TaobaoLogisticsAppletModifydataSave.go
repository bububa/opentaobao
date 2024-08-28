package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsAppletModifydataSave 物流小程序修改物流信息回传接口
// taobao.logistics.applet.modifydata.save
//
// 物流小程序修改物流信息回传接口
func TaobaoLogisticsAppletModifydataSave(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoLogisticsAppletModifydataSaveAPIRequest, resp *mtopopen.TaobaoLogisticsAppletModifydataSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
