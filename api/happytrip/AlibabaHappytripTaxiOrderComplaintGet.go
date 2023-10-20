package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiordercomplaintget 投诉详情
// alibaba.happytrip.taxi.order.complaint.get
//
// 获取投诉处理进度详情
func Alibabahappytriptaxiordercomplaintget(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiordercomplaintgetAPIRequest, session string) (*happytrip.AlibabahappytriptaxiordercomplaintgetAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiordercomplaintgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
