package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountIndexFindAPIRequest
Open Account索引查询 API请求
taobao.open.account.index.find

Open Account索引查询 */
type TaobaoOpenAccountIndexFindAPIRequest struct {
	model.Params
	// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
	_indexType int64
	// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
	_indexValue string
}

// New
