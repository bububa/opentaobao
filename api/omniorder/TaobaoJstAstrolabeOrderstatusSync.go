package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaojstastrolabeorderstatussync 线下门店派单以及单据相关操作接口
// taobao.jst.astrolabe.orderstatus.sync
//
// 针对ERP系统部署在门店的商家，将派单状态回流到星盘
func Taobaojstastrolabeorderstatussync(clt *core.SDKClient, req *omniorder.TaobaojstastrolabeorderstatussyncAPIRequest, session string) (*omniorder.TaobaojstastrolabeorderstatussyncAPIResponse, error) {
	var resp omniorder.TaobaojstastrolabeorderstatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
