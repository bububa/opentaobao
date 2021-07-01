package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildRootnodeGetAPIRequest
获取少儿大厅根类目接口 API请求
yunos.tvpubadmin.content.child.rootnode.get

通过此接口可获取少儿大厅根类目列表 */
type YunosTvpubadminContentChildRootnodeGetAPIRequest struct {
	model.Params
	// 是否需要首页目录
	_needHomePage bool
}

// New
