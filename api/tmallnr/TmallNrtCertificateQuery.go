package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCertificateQuery 批量查询电子凭证信息
// tmall.nrt.certificate.query
//
// 批量查询电子凭证信息
func TmallNrtCertificateQuery(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrtCertificateQueryAPIRequest, resp *tmallnr.TmallNrtCertificateQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
