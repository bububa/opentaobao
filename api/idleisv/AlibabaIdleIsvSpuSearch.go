package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvspusearch spu搜索接口
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
func Alibabaidleisvspusearch(clt *core.SDKClient, req *idleisv.AlibabaidleisvspusearchAPIRequest, session string) (*idleisv.AlibabaidleisvspusearchAPIResponse, error) {
	var resp idleisv.AlibabaidleisvspusearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
