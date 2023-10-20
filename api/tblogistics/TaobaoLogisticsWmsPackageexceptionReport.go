package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticswmspackageexceptionreport 无主件回告
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
func Taobaologisticswmspackageexceptionreport(clt *core.SDKClient, req *tblogistics.TaobaologisticswmspackageexceptionreportAPIRequest, session string) (*tblogistics.TaobaologisticswmspackageexceptionreportAPIResponse, error) {
	var resp tblogistics.TaobaologisticswmspackageexceptionreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
