package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaoweikesubscinfoget 需求订单查询接口
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
func Taobaoweikesubscinfoget(clt *core.SDKClient, req *servicecenter.TaobaoweikesubscinfogetAPIRequest, session string) (*servicecenter.TaobaoweikesubscinfogetAPIResponse, error) {
	var resp servicecenter.TaobaoweikesubscinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
