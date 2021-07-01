package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminManageTopicContentadd
专题新增内容
yunos.tvpubadmin.manage.topic.contentadd

专题新增内容 */
func YunosTvpubadminManageTopicContentadd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContentaddAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicContentaddAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminManageTopicContentaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
