package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemOptionPageAPIRequest
分页查询定向标签列表 API请求
taobao.feedflow.item.option.page

分页查询定向标签列表 */
type TaobaoFeedflowItemOptionPageAPIRequest struct {
	model.Params
	// 标签查询条件
	_labelQuery *LabelQueryDto
}

// New
