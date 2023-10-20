package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarxcarsynchronizecarmodeldata 爱车车型数据同步
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
func Tmallcarxcarsynchronizecarmodeldata(clt *core.SDKClient, req *tmallcar.TmallcarxcarsynchronizecarmodeldataAPIRequest, session string) (*tmallcar.TmallcarxcarsynchronizecarmodeldataAPIResponse, error) {
	var resp tmallcar.TmallcarxcarsynchronizecarmodeldataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
