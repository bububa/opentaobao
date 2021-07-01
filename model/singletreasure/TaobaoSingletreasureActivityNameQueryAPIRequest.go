package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityNameQueryAPIRequest
查询官方的活动名称接口 API请求
taobao.singletreasure.activity.name.query

查询官方的活动名称列表接口 */
type TaobaoSingletreasureActivityNameQueryAPIRequest struct {
	model.Params
}

// New
