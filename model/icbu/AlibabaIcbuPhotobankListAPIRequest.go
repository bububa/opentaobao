package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuPhotobankListAPIRequest
国际站图片银行查询接口 API请求
alibaba.icbu.photobank.list

国际站图片银行查询接口 */
type AlibabaIcbuPhotobankListAPIRequest struct {
	model.Params
	// 当前翻页数
	_currentPage int64
	// 图片组id
	_groupId string
	// 存放位置 必要条件, 包括ALL_GROUP(所有目录), SUB_GROUP(指定图片组下),UNGROUP(未分组), TEMP(disable)四个值
	_locationType string
	// 每页显示数
	_pageSize int64
	// 额外的上下文信息. 例如:icvId
	_extraContext string
}

// New
