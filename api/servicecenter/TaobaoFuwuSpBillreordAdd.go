package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoFuwuSpBillreordAdd 内购服务确认单明细上传接口
// taobao.fuwu.sp.billreord.add
//
// isv能通过该接口上传确认单明细数据
func TaobaoFuwuSpBillreordAdd(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoFuwuSpBillreordAddAPIRequest, resp *servicecenter.TaobaoFuwuSpBillreordAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
