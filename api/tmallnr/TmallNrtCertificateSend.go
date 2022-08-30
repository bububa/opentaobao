package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCertificateSend 有价礼包发放电子凭证
// tmall.nrt.certificate.send
//
// 支持有价礼包发放电子凭证
func TmallNrtCertificateSend(clt *core.SDKClient, req *tmallnr.TmallNrtCertificateSendAPIRequest, session string) (*tmallnr.TmallNrtCertificateSendAPIResponse, error) {
	var resp tmallnr.TmallNrtCertificateSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
