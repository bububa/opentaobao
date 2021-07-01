package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowGetbyshowlongidAPIRequest
迎客松根据节目longid获取节目元数据 API请求
yunos.tvpubadmin.content.show.getbyshowlongid

迎客松根据节目longid获取节目元数据 */
type YunosTvpubadminContentShowGetbyshowlongidAPIRequest struct {
	model.Params
	// 节目longid
	_showLongId int64
}

// New
