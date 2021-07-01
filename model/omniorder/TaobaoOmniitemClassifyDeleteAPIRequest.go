package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyDeleteAPIRequest
删除一个分类 API请求
taobao.omniitem.classify.delete

删除一个分类 */
type TaobaoOmniitemClassifyDeleteAPIRequest struct {
	model.Params
	// 分类ID
	_classifyId int64
	// 操作人信息（暂时不填）
	_operator string
}

// New
