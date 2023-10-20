package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomordercheckorderinfo 金融购机订单信息校验
// alibaba.alicom.order.checkorderinfo
//
// 金融购机订单信息校验
func Alibabaalicomordercheckorderinfo(clt *core.SDKClient, req *alicom.AlibabaalicomordercheckorderinfoAPIRequest, session string) (*alicom.AlibabaalicomordercheckorderinfoAPIResponse, error) {
	var resp alicom.AlibabaalicomordercheckorderinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
