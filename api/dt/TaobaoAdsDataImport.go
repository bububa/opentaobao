package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// Taobaoadsdataimport 数据导入
// taobao.ads.data.import
//
// 数据导入
func Taobaoadsdataimport(clt *core.SDKClient, req *dt.TaobaoadsdataimportAPIRequest, session string) (*dt.TaobaoadsdataimportAPIResponse, error) {
	var resp dt.TaobaoadsdataimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
