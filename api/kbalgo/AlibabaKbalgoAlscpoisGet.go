package kbalgo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kbalgo"
)

// AlibabaKbalgoAlscpoisGet 百度批量获取本地poi接口
// alibaba.kbalgo.alscpois.get
//
// 接口用于百度方获取本地生活poi数据，分页获取。
func AlibabaKbalgoAlscpoisGet(ctx context.Context, clt *core.SDKClient, req *kbalgo.AlibabaKbalgoAlscpoisGetAPIRequest, resp *kbalgo.AlibabaKbalgoAlscpoisGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
