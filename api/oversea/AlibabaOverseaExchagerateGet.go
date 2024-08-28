package oversea

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/oversea"
)

// AlibabaOverseaExchagerateGet 汇率信息获取
// alibaba.oversea.exchagerate.get
//
// 提供外部汇率查询接口
func AlibabaOverseaExchagerateGet(ctx context.Context, clt *core.SDKClient, req *oversea.AlibabaOverseaExchagerateGetAPIRequest, resp *oversea.AlibabaOverseaExchagerateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
