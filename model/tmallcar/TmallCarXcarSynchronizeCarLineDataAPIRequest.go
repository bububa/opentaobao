package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarXcarSynchronizeCarLineDataAPIRequest
我的爱卡车型配置数据 API请求
tmall.car.xcar.synchronize.car.line.data

同步我的爱卡车系配置数据到天猫汽车 */
type TmallCarXcarSynchronizeCarLineDataAPIRequest struct {
	model.Params
	// 入参对象
	_paramXCarSysLineDTO *XCarSysLineDto
}

// New
