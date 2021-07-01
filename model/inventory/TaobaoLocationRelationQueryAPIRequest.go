package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLocationRelationQueryAPIRequest
地点关联关系查询 API请求
taobao.location.relation.query

地点关联关系查询
门店和仓库关联关系查询 */
type TaobaoLocationRelationQueryAPIRequest struct {
	model.Params
	// 关系查询
	_locationRelation *LocationRelationDto
}

// New
