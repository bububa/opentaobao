package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaopaimainftcertificateapplycallback 数字藏品版权证书申请结果回调
// taobao.paimai.nft.certificate.applycallback
//
// 数字藏品版权证书申请结果回调
func Taobaopaimainftcertificateapplycallback(clt *core.SDKClient, req *paimai.TaobaopaimainftcertificateapplycallbackAPIRequest, session string) (*paimai.TaobaopaimainftcertificateapplycallbackAPIResponse, error) {
	var resp paimai.TaobaopaimainftcertificateapplycallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
