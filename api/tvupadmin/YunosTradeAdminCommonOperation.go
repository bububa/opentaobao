package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostradeadmincommonoperation 交易迎客松通用服务接口
// yunos.trade.admin.common.operation
//
// 迎客松交易相关通用接口
func Yunostradeadmincommonoperation(clt *core.SDKClient, req *tvupadmin.YunostradeadmincommonoperationAPIRequest, session string) (*tvupadmin.YunostradeadmincommonoperationAPIResponse, error) {
	var resp tvupadmin.YunostradeadmincommonoperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
