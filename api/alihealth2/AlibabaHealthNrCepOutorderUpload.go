package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaHealthNrCepOutorderUpload 线上订单收货验收单、出入库单据生成接口
// alibaba.health.nr.cep.outorder.upload
//
// 线上订单收货验收单、出入库单据生成接口
func AlibabaHealthNrCepOutorderUpload(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaHealthNrCepOutorderUploadAPIRequest, resp *alihealth2.AlibabaHealthNrCepOutorderUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
