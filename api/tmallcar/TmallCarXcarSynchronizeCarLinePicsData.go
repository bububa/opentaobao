package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarxcarsynchronizecarlinepicsdata 爱卡车系图片数据接入
// tmall.car.xcar.synchronize.car.line.pics.data
//
// 爱卡车系图片数据同步天猫汽车
func Tmallcarxcarsynchronizecarlinepicsdata(clt *core.SDKClient, req *tmallcar.TmallcarxcarsynchronizecarlinepicsdataAPIRequest, session string) (*tmallcar.TmallcarxcarsynchronizecarlinepicsdataAPIResponse, error) {
	var resp tmallcar.TmallcarxcarsynchronizecarlinepicsdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
