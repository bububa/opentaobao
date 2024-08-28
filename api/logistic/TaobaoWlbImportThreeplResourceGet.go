package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoWlbImportThreeplResourceGet 3PL直邮获取资源列表
// taobao.wlb.import.threepl.resource.get
//
// 获取3pl直邮的发货可用资源
func TaobaoWlbImportThreeplResourceGet(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoWlbImportThreeplResourceGetAPIRequest, resp *logistic.TaobaoWlbImportThreeplResourceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
