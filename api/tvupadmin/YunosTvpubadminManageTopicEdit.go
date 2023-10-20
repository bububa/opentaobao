package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopicedit 编辑专题
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
func Yunostvpubadminmanagetopicedit(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopiceditAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopiceditAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopiceditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
