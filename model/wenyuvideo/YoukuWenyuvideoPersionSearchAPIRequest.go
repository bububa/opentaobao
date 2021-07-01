package wenyuvideo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuWenyuvideoPersionSearchAPIRequest
根据人物名称查询人物列表 API请求
youku.wenyuvideo.persion.search

根据人物名称查询人物列表 */
type YoukuWenyuvideoPersionSearchAPIRequest struct {
	model.Params
	// 人物名字，搜索规则是完全匹配，即只返回同名人物列表
	_personName string
}

// New
