package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarxcarsynchronizecarlinedata 我的爱卡车型配置数据
// tmall.car.xcar.synchronize.car.line.data
//
// 同步我的爱卡车系配置数据到天猫汽车
func Tmallcarxcarsynchronizecarlinedata(clt *core.SDKClient, req *tmallcar.TmallcarxcarsynchronizecarlinedataAPIRequest, session string) (*tmallcar.TmallcarxcarsynchronizecarlinedataAPIResponse, error) {
	var resp tmallcar.TmallcarxcarsynchronizecarlinedataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
