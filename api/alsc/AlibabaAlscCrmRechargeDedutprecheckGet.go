package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargededutprecheckget 储值核销预先校验
// alibaba.alsc.crm.recharge.dedutprecheck.get
//
// 储值核销预先校验接口
func Alibabaalsccrmrechargededutprecheckget(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargededutprecheckgetAPIRequest, session string) (*alsc.AlibabaalsccrmrechargededutprecheckgetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargededutprecheckgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
