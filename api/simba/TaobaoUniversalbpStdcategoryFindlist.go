package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpstdcategoryfindlist 人群相关类目查询
// taobao.universalbp.stdcategory.findlist
//
// 入参商品信息，出参商品所属类别
func Taobaouniversalbpstdcategoryfindlist(clt *core.SDKClient, req *simba.TaobaouniversalbpstdcategoryfindlistAPIRequest, session string) (*simba.TaobaouniversalbpstdcategoryfindlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpstdcategoryfindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
