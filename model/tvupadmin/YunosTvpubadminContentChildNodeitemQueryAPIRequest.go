package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildNodeitemQueryAPIRequest
查询少儿大厅类目内容 API请求
yunos.tvpubadmin.content.child.nodeitem.query

查询少儿大厅类目内容信息 */
type YunosTvpubadminContentChildNodeitemQueryAPIRequest struct {
	model.Params
	// 主键ID
	_id int64
	// 类目ID
	_nodeId int64
	// 状态
	_status int64
	// 页码
	_pageNo int64
	// 内容名称
	_name string
	// 单页数量
	_pageSize int64
}

// New
