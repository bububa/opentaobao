package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbareportcityget 获取城市维度报表
// taobao.simba.report.city.get
//
// 获取城市维度报表
func Taobaosimbareportcityget(clt *core.SDKClient, req *simba.TaobaosimbareportcitygetAPIRequest, session string) (*simba.TaobaosimbareportcitygetAPIResponse, error) {
	var resp simba.TaobaosimbareportcitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
