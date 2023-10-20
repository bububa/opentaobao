package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargededutupdate 储值消费
// alibaba.alsc.crm.recharge.dedut.update
//
// 增加储值消费接口
func Alibabaalsccrmrechargededutupdate(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargededutupdateAPIRequest, session string) (*alsc.AlibabaalsccrmrechargededutupdateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargededutupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
