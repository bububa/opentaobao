package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptravelsync 差旅申请单同步接口
// alibaba.happytrip.travel.sync
//
// 以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
func Alibabahappytriptravelsync(clt *core.SDKClient, req *happytrip.AlibabahappytriptravelsyncAPIRequest, session string) (*happytrip.AlibabahappytriptravelsyncAPIResponse, error) {
	var resp happytrip.AlibabahappytriptravelsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
