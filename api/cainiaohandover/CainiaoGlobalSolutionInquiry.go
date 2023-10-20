package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalsolutioninquiry 解决方案询盘
// cainiao.global.solution.inquiry
//
// 根据交易单号查询可用的解决方案
func Cainiaoglobalsolutioninquiry(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalsolutioninquiryAPIRequest, session string) (*cainiaohandover.CainiaoglobalsolutioninquiryAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalsolutioninquiryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
