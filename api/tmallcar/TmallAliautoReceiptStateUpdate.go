package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoreceiptstateupdate 服务工单状态更新
// tmall.aliauto.receipt.state.update
//
// 二轮车服务工单状态更新
func Tmallaliautoreceiptstateupdate(clt *core.SDKClient, req *tmallcar.TmallaliautoreceiptstateupdateAPIRequest, session string) (*tmallcar.TmallaliautoreceiptstateupdateAPIResponse, error) {
	var resp tmallcar.TmallaliautoreceiptstateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
