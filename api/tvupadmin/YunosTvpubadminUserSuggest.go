package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminusersuggest 获取关联账户列表
// yunos.tvpubadmin.user.suggest
//
// 获取关联账户列表
func Yunostvpubadminusersuggest(clt *core.SDKClient, req *tvupadmin.YunostvpubadminusersuggestAPIRequest, session string) (*tvupadmin.YunostvpubadminusersuggestAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminusersuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
