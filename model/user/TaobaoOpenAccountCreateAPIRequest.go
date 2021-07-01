package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountCreateAPIRequest
Open Account导入数据 API请求
taobao.open.account.create

Open Account导入数据 */
type TaobaoOpenAccountCreateAPIRequest struct {
	model.Params
	// Open Account的列表
	_paramList []OpenAccount
}

// New
