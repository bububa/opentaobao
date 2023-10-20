package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopicfindbyid 根据id获取专题信息
// yunos.tvpubadmin.manage.topic.findbyid
//
// 根据id获取专题信息
func Yunostvpubadminmanagetopicfindbyid(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopicfindbyidAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopicfindbyidAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopicfindbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
