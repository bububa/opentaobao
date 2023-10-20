package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// Taobaoadsdataquery 导入数据查询
// taobao.ads.data.query
//
// 导入数据查询
func Taobaoadsdataquery(clt *core.SDKClient, req *dt.TaobaoadsdataqueryAPIRequest, session string) (*dt.TaobaoadsdataqueryAPIResponse, error) {
	var resp dt.TaobaoadsdataqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
