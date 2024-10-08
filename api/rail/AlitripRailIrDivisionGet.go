package rail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailIrDivisionGet 国际火车票标准城市查询
// alitrip.rail.ir.division.get
//
// 国际火车票提供给代理商用于查询标准城市信息，全部城市数据量209530条，含除中国大陆以外的全部海外区域。
// 代理商通过分页查询的方式，拉取飞猪平台方全部境外标准城市，用于自身城市与飞猪平台城市的映射。
func AlitripRailIrDivisionGet(ctx context.Context, clt *core.SDKClient, req *rail.AlitripRailIrDivisionGetAPIRequest, resp *rail.AlitripRailIrDivisionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
