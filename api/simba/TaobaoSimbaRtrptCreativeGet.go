package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbartrptcreativeget 获取创意实时报表数据
// taobao.simba.rtrpt.creative.get
//
// 获取创意实时报表数据
func Taobaosimbartrptcreativeget(clt *core.SDKClient, req *simba.TaobaosimbartrptcreativegetAPIRequest, session string) (*simba.TaobaosimbartrptcreativegetAPIResponse, error) {
	var resp simba.TaobaosimbartrptcreativegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
