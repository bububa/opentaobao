package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverPdfGet 获取面单PDF文件数据
// cainiao.global.handover.pdf.get
//
// 返回指定大包面单的PDF文件数据
func CainiaoGlobalHandoverPdfGet(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverPdfGetAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverPdfGetAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalHandoverPdfGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
