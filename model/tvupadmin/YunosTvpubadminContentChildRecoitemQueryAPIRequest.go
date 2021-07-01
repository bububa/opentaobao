package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildRecoitemQueryAPIRequest
查询少儿大厅推荐内容列表 API请求
yunos.tvpubadmin.content.child.recoitem.query

查询少儿大厅推荐内容列表 */
type YunosTvpubadminContentChildRecoitemQueryAPIRequest struct {
	model.Params
	// 主键ID
	_id int64
	// 所属类目ID
	_nodeId int64
	// 状态
	_status int64
	// 页码
	_pageNo int64
	// 名称
	_name string
	// 单页数量
	_pageSize int64
}

// New
