package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Alibabamosfalconposcounterquery 云POS查看专柜属性
// alibaba.mos.falcon.pos.counter.query
//
// 银泰商业获取专柜是否支持小数等属性查看
func Alibabamosfalconposcounterquery(clt *core.SDKClient, req *util.AlibabamosfalconposcounterqueryAPIRequest, session string) (*util.AlibabamosfalconposcounterqueryAPIResponse, error) {
	var resp util.AlibabamosfalconposcounterqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
