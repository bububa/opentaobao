package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrtcertificatequery 批量查询电子凭证信息
// tmall.nrt.certificate.query
//
// 批量查询电子凭证信息
func Tmallnrtcertificatequery(clt *core.SDKClient, req *tmallnr.TmallnrtcertificatequeryAPIRequest, session string) (*tmallnr.TmallnrtcertificatequeryAPIResponse, error) {
	var resp tmallnr.TmallnrtcertificatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
