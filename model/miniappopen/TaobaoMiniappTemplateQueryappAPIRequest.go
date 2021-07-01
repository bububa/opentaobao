package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateQueryappAPIRequest
查询实例化应用版本 API请求
taobao.miniapp.template.queryapp

根据模板id和商家信息，查询实例化小程序版本查询 */
type TaobaoMiniappTemplateQueryappAPIRequest struct {
	model.Params
	// 分页大小，最大50，按照小程序Id倒序
	_pageSize int64
	// 模板id
	_templateId string
	// 分页号,>=1
	_pageNum int64
}

// New
