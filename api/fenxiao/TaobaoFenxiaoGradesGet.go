package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoGradesGet 分销商等级查询
// taobao.fenxiao.grades.get
//
// 根据供应商ID，查询他的分销商等级信息
func TaobaoFenxiaoGradesGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoGradesGetAPIRequest, resp *fenxiao.TaobaoFenxiaoGradesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
