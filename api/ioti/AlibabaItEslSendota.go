package ioti

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItEslSendota 电子价签ota接口
// alibaba.it.esl.sendota
//
// 厂测接口，电子价签ota接口
func AlibabaItEslSendota(ctx context.Context, clt *core.SDKClient, req *ioti.AlibabaItEslSendotaAPIRequest, resp *ioti.AlibabaItEslSendotaAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
