package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundcreate 创建OpenMall退款单
// taobao.openmall.refund.create
//
// 创建OpenMall退款单
// 如存在未完结的退款单，则返回该退款单ID
func Taobaoopenmallrefundcreate(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundcreateAPIRequest, session string) (*openmall.TaobaoopenmallrefundcreateAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
