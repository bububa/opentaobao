package ioti

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItEslSendled 厂测LED控制
// alibaba.it.esl.sendled
//
// 针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
func AlibabaItEslSendled(ctx context.Context, clt *core.SDKClient, req *ioti.AlibabaItEslSendledAPIRequest, resp *ioti.AlibabaItEslSendledAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
