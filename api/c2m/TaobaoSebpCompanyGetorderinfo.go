package c2m

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/c2m"
)

// TaobaoSebpCompanyGetorderinfo 淘小铺公司订单信息
// taobao.sebp.company.getorderinfo
//
// 淘小铺合作公司获取公司订单信息，用于公司结算使用
func TaobaoSebpCompanyGetorderinfo(clt *core.SDKClient, req *c2m.TaobaoSebpCompanyGetorderinfoAPIRequest, session string) (*c2m.TaobaoSebpCompanyGetorderinfoAPIResponse, error) {
	var resp c2m.TaobaoSebpCompanyGetorderinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
