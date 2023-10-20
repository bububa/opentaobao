package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopicadd 新增专题
// yunos.tvpubadmin.manage.topic.add
//
// 专题新增
func Yunostvpubadminmanagetopicadd(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopicaddAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopicaddAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopicaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
