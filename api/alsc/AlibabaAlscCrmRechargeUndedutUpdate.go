package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeundedutupdate 储值消费退款(逆向)
// alibaba.alsc.crm.recharge.undedut.update
//
// 新增储值消费退款接口
func Alibabaalsccrmrechargeundedutupdate(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeundedutupdateAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeundedutupdateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeundedutupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
