package nrpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest
去前置机商品在线查询 API请求
alibaba.mos.commdy.posmerchandise.getmerchandise

去前置机商品在线查询接口 */
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest struct {
	model.Params
	// 查询参数列表
	_posMerchandiseList []QueryMerchandiseDto
}

// New
