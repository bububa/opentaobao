package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbCrossborderWaybillGet 集货商家pdf和云打印面单获取，pdf需要配置白名单
// taobao.wlb.crossborder.waybill.get
//
// 【TOF】欧洲供应商PDF格式电子面单渲染下发
//
//	需求链接：https://aone.alibaba-inc.com/req/21210808
func TaobaoWlbCrossborderWaybillGet(ctx context.Context, clt *core.SDKClient, req *wlbimports.TaobaoWlbCrossborderWaybillGetAPIRequest, resp *wlbimports.TaobaoWlbCrossborderWaybillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
