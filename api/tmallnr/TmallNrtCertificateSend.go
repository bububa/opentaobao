package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrtcertificatesend 有价礼包发放电子凭证
// tmall.nrt.certificate.send
//
// 支持有价礼包发放电子凭证
func Tmallnrtcertificatesend(clt *core.SDKClient, req *tmallnr.TmallnrtcertificatesendAPIRequest, session string) (*tmallnr.TmallnrtcertificatesendAPIResponse, error) {
	var resp tmallnr.TmallnrtcertificatesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
