package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCertificateQuery 批量查询电子凭证信息
// tmall.nrt.certificate.query
//
// 批量查询电子凭证信息
func TmallNrtCertificateQuery(clt *core.SDKClient, req *tmallnr.TmallNrtCertificateQueryAPIRequest, resp *tmallnr.TmallNrtCertificateQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
