package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbartrptbidwordget 获取推广词实时报表数据
// taobao.simba.rtrpt.bidword.get
//
// 获取推广词报表数据
func Taobaosimbartrptbidwordget(clt *core.SDKClient, req *simba.TaobaosimbartrptbidwordgetAPIRequest, session string) (*simba.TaobaosimbartrptbidwordgetAPIResponse, error) {
	var resp simba.TaobaosimbartrptbidwordgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
