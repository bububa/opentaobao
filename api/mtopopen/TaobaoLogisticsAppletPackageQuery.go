package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaologisticsappletpackagequery 淘宝包裹查询
// taobao.logistics.applet.package.query
//
// 淘宝包裹查询
func Taobaologisticsappletpackagequery(clt *core.SDKClient, req *mtopopen.TaobaologisticsappletpackagequeryAPIRequest, session string) (*mtopopen.TaobaologisticsappletpackagequeryAPIResponse, error) {
	var resp mtopopen.TaobaologisticsappletpackagequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
