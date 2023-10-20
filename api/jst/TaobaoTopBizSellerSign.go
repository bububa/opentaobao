package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaotopbizsellersign 淘宝订单履约-商家erp签约
// taobao.top.biz.seller.sign
//
// 淘宝订单履约-商家erp签约，包含各场景的签约
func Taobaotopbizsellersign(clt *core.SDKClient, req *jst.TaobaotopbizsellersignAPIRequest, session string) (*jst.TaobaotopbizsellersignAPIResponse, error) {
	var resp jst.TaobaotopbizsellersignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
