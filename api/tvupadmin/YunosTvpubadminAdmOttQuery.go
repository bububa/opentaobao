package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminadmottquery 优酷OTT端广告素材查询
// yunos.tvpubadmin.adm.ott.query
//
// 查询广告素材
func Yunostvpubadminadmottquery(clt *core.SDKClient, req *tvupadmin.YunostvpubadminadmottqueryAPIRequest, session string) (*tvupadmin.YunostvpubadminadmottqueryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminadmottqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
