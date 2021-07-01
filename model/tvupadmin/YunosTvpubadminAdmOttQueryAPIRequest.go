package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminAdmOttQueryAPIRequest
优酷OTT端广告素材查询 API请求
yunos.tvpubadmin.adm.ott.query

查询广告素材 */
type YunosTvpubadminAdmOttQueryAPIRequest struct {
	model.Params
	// 查询参数json格式
	_query string
}

// New
