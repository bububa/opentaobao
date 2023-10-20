package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomvtopentradegetproductinfo 查询新虚拟产品配置信息
// alibaba.alicom.vt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
func Alibabaalicomvtopentradegetproductinfo(clt *core.SDKClient, req *alicom.AlibabaalicomvtopentradegetproductinfoAPIRequest, session string) (*alicom.AlibabaalicomvtopentradegetproductinfoAPIResponse, error) {
	var resp alicom.AlibabaalicomvtopentradegetproductinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
