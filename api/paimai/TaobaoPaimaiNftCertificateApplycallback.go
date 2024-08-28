package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiNftCertificateApplycallback 数字藏品版权证书申请结果回调
// taobao.paimai.nft.certificate.applycallback
//
// 数字藏品版权证书申请结果回调
func TaobaoPaimaiNftCertificateApplycallback(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoPaimaiNftCertificateApplycallbackAPIRequest, resp *paimai.TaobaoPaimaiNftCertificateApplycallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
