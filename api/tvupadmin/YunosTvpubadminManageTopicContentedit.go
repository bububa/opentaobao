package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicContentedit 专题关联内容编辑
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
func YunosTvpubadminManageTopicContentedit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContenteditAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicContenteditAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminManageTopicContenteditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
