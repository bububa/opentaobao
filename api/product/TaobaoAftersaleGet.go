package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaoaftersaleget 查询用户售后服务模板
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
func Taobaoaftersaleget(clt *core.SDKClient, req *product.TaobaoaftersalegetAPIRequest, session string) (*product.TaobaoaftersalegetAPIResponse, error) {
	var resp product.TaobaoaftersalegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
