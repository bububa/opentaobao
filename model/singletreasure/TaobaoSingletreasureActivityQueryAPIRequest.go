package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityQueryAPIRequest
查询活动列表接口 API请求
taobao.singletreasure.activity.query

查询活动列表接口 */
type TaobaoSingletreasureActivityQueryAPIRequest struct {
	model.Params
	// 查询对象
	_query *PageQueryDto
}

// New
