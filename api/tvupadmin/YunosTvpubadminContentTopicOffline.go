package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontenttopicoffline 迎客松专题下线
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
func Yunostvpubadmincontenttopicoffline(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontenttopicofflineAPIRequest, session string) (*tvupadmin.YunostvpubadmincontenttopicofflineAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontenttopicofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
