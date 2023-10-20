package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbahourreportadgroupget 推广单元小时级别实时报表查询
// taobao.simba.hour.report.adgroup.get
//
// 推广单元小时级别实时报表查询
func Taobaosimbahourreportadgroupget(clt *core.SDKClient, req *simba.TaobaosimbahourreportadgroupgetAPIRequest, session string) (*simba.TaobaosimbahourreportadgroupgetAPIResponse, error) {
	var resp simba.TaobaosimbahourreportadgroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
