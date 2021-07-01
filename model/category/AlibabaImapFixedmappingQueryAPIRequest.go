package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaImapFixedmappingQueryAPIRequest
查询两个渠道之间的固定映射关系，不通过算法兜底 API请求
alibaba.imap.fixedmapping.query

查询两个渠道之间的固定映射关系，不通过算法兜底 */
type AlibabaImapFixedmappingQueryAPIRequest struct {
	model.Params
	// 密码
	_password string
	// 账号
	_appName string
	// 源渠道ID
	_srcChannelId int64
	// 目标渠道ID列表
	_targetChannelIdList []int64
	// 目标渠道ID
	_targetCategoryId int64
	// 源类目ID
	_srcCategoryId int64
}

// New
