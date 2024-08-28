package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverPdfGet 获取面单PDF文件数据
// cainiao.global.handover.pdf.get
//
// 返回指定大包面单的PDF文件数据
func CainiaoGlobalHandoverPdfGet(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverPdfGetAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverPdfGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
