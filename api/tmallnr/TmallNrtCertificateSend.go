package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCertificateSend 有价礼包发放电子凭证
// tmall.nrt.certificate.send
//
// 支持有价礼包发放电子凭证
func TmallNrtCertificateSend(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrtCertificateSendAPIRequest, resp *tmallnr.TmallNrtCertificateSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
