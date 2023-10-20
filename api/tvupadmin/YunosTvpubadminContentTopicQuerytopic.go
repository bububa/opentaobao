package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontenttopicquerytopic 迎客松专题查询
// yunos.tvpubadmin.content.topic.querytopic
//
// 迎客松专题查询
func Yunostvpubadmincontenttopicquerytopic(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontenttopicquerytopicAPIRequest, session string) (*tvupadmin.YunostvpubadmincontenttopicquerytopicAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontenttopicquerytopicAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
