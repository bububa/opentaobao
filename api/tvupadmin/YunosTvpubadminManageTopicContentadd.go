package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopiccontentadd 专题新增内容
// yunos.tvpubadmin.manage.topic.contentadd
//
// 专题新增内容
func Yunostvpubadminmanagetopiccontentadd(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopiccontentaddAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopiccontentaddAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopiccontentaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
