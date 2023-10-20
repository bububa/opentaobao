package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcreativepreadd 创建单品创意前置信息
// taobao.universalbp.creative.preadd
//
// 用于关键词场景创建单品创意前使用
func Taobaouniversalbpcreativepreadd(clt *core.SDKClient, req *simba.TaobaouniversalbpcreativepreaddAPIRequest, session string) (*simba.TaobaouniversalbpcreativepreaddAPIResponse, error) {
	var resp simba.TaobaouniversalbpcreativepreaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
