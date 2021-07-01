package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowGetbyshowidAPIRequest
迎客松根据节目id获取节目元数据 API请求
yunos.tvpubadmin.content.show.getbyshowid

迎客松根据节目id获取节目元数据 */
type YunosTvpubadminContentShowGetbyshowidAPIRequest struct {
	model.Params
	// 节目字符串id
	_showId string
}

// New
