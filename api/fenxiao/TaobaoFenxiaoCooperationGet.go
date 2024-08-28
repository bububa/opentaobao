package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoCooperationGet 供应商或分销商获取合作关系信息
// taobao.fenxiao.cooperation.get
//
// 获取供应商的合作关系信息
func TaobaoFenxiaoCooperationGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoCooperationGetAPIRequest, resp *fenxiao.TaobaoFenxiaoCooperationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
