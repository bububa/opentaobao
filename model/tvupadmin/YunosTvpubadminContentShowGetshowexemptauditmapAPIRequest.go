package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest
迎客松批量查询节目某个牌照的免审状态 API请求
yunos.tvpubadmin.content.show.getshowexemptauditmap

迎客松批量查询节目某个牌照的免审状态 */
type YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest struct {
	model.Params
	// 节目longid
	_showLongIds string
	// 牌照id
	_license int64
}

// New
