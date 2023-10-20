package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoscitemoutercodeget 根据outerCode查询商品
// taobao.scitem.outercode.get
//
// 根据outerCode查询商品
func Taobaoscitemoutercodeget(clt *core.SDKClient, req *fenxiao.TaobaoscitemoutercodegetAPIRequest, session string) (*fenxiao.TaobaoscitemoutercodegetAPIResponse, error) {
	var resp fenxiao.TaobaoscitemoutercodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
