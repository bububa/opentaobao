package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScUcrowdReportGet 淘宝客-服务商-人群推广效果
// taobao.tbk.sc.ucrowd.report.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群的推广和转化效果数据。
func TaobaoTbkScUcrowdReportGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScUcrowdReportGetAPIRequest, resp *tbk.TaobaoTbkScUcrowdReportGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
