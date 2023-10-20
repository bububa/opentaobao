package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTopicQuerytopic 迎客松专题查询
// yunos.tvpubadmin.content.topic.querytopic
//
// 迎客松专题查询
func YunosTvpubadminContentTopicQuerytopic(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTopicQuerytopicAPIRequest, resp *tvupadmin.YunosTvpubadminContentTopicQuerytopicAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
