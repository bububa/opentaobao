package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTicketScenicQueryAPIRequest
【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） API请求
alitrip.ticket.scenic.query

查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！ */
type AlitripTicketScenicQueryAPIRequest struct {
	model.Params
	// 标准景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
	_aliScenicId int64
	// 当前分页。每页默认最多返回20条数据
	_currentPage int64
	// 商家景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
	_outScenicId string
}

// New
