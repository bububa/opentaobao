package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiNftCertificateApplycallback 数字藏品版权证书申请结果回调
// taobao.paimai.nft.certificate.applycallback
//
// 数字藏品版权证书申请结果回调
func TaobaoPaimaiNftCertificateApplycallback(clt *core.SDKClient, req *paimai.TaobaoPaimaiNftCertificateApplycallbackAPIRequest, session string) (*paimai.TaobaoPaimaiNftCertificateApplycallbackAPIResponse, error) {
	var resp paimai.TaobaoPaimaiNftCertificateApplycallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
