package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentVideoGetauditlistAPIRequest
迎客松视频审核记录查询 API请求
yunos.tvpubadmin.content.video.getauditlist

迎客松视频审核记录查询 */
type YunosTvpubadminContentVideoGetauditlistAPIRequest struct {
	model.Params
	// 查询
	_query string
}

// New
