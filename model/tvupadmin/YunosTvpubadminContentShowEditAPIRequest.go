package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowEditAPIRequest
媒资节目信息修改 API请求
yunos.tvpubadmin.content.show.edit

供迎客松修改媒资节目信息 */
type YunosTvpubadminContentShowEditAPIRequest struct {
	model.Params
	// 请求入参，JSON格式
	_data string
}

// New
