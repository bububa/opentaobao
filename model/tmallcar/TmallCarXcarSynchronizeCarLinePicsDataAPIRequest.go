package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarXcarSynchronizeCarLinePicsDataAPIRequest
爱卡车系图片数据接入 API请求
tmall.car.xcar.synchronize.car.line.pics.data

爱卡车系图片数据同步天猫汽车 */
type TmallCarXcarSynchronizeCarLinePicsDataAPIRequest struct {
	model.Params
	// 入参对象
	_paramXCarSysLinePicsDTO *XCarSysLinePicsDto
}

// New
