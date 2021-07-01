package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelGetEntityByTagAPIRequest
根据标签查询实体 API请求
taobao.xhotel.get.entity.by.tag

根据标签查询实体 */
type TaobaoXhotelGetEntityByTagAPIRequest struct {
	model.Params
	// 标签
	_tag int64
	// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
	_tokenStr string
}

// New
