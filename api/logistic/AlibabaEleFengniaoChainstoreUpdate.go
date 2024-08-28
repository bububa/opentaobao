package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoChainstoreUpdate 修改门店信息接口
// alibaba.ele.fengniao.chainstore.update
//
// 修改门店的经纬度，文本地址，电话，门店名
func AlibabaEleFengniaoChainstoreUpdate(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreUpdateAPIRequest, resp *logistic.AlibabaEleFengniaoChainstoreUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
