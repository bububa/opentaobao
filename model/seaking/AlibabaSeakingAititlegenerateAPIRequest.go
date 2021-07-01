package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingAititlegenerateAPIRequest
标题智能优化 API请求
alibaba.seaking.aititlegenerate

标题智能优化 */
type AlibabaSeakingAititlegenerateAPIRequest struct {
	model.Params
	// erp用户id
	_identifier string
	// 扩展信息
	_extra *Extra
	// 语种
	_language string
	// 商品属性
	_attributes string
	// 调用来源(erp名称)
	_identifierType string
	// 标题
	_title string
	// 商品所在平台（ae/icbu）
	_platform string
	// 类目id,没有的时候传-1
	_categoryId int64
}

// New
