package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarXcarSynchronizeCarModelDataAPIRequest
爱车车型数据同步 API请求
tmall.car.xcar.synchronize.car.model.data

爱车汽车车型数据同步到天猫 */
type TmallCarXcarSynchronizeCarModelDataAPIRequest struct {
	model.Params
	// 传入对象描述
	_paramXCarSysModelDTO *XCarSysModelDto
}

// New
