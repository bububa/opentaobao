package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarXcarSynchronizeCarModelData 爱车车型数据同步
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
func TmallCarXcarSynchronizeCarModelData(clt *core.SDKClient, req *tmallcar.TmallCarXcarSynchronizeCarModelDataAPIRequest, resp *tmallcar.TmallCarXcarSynchronizeCarModelDataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
