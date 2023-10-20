package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomwttopentradegetproductinfo 查询存送产品信息
// alibaba.alicom.wtt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
func Alibabaalicomwttopentradegetproductinfo(clt *core.SDKClient, req *alicom.AlibabaalicomwttopentradegetproductinfoAPIRequest, session string) (*alicom.AlibabaalicomwttopentradegetproductinfoAPIResponse, error) {
	var resp alicom.AlibabaalicomwttopentradegetproductinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
