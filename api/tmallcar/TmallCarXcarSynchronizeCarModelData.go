package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarXcarSynchronizeCarModelData 爱车车型数据同步
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
func TmallCarXcarSynchronizeCarModelData(clt *core.SDKClient, req *tmallcar.TmallCarXcarSynchronizeCarModelDataAPIRequest, session string) (*tmallcar.TmallCarXcarSynchronizeCarModelDataAPIResponse, error) {
	var resp tmallcar.TmallCarXcarSynchronizeCarModelDataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
