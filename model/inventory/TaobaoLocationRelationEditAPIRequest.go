package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLocationRelationEditAPIRequest
地点关联关系增量编辑 API请求
taobao.location.relation.edit

地点关联关系增量编辑 */
type TaobaoLocationRelationEditAPIRequest struct {
	model.Params
	// 关系对象列表
	_locationRelationList []LocationRelationDto
}

// New
