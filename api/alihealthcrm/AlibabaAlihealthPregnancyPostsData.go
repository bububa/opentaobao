package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyPostsData 发回帖子信息同步
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
func AlibabaAlihealthPregnancyPostsData(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyPostsDataAPIRequest, resp *alihealthcrm.AlibabaAlihealthPregnancyPostsDataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
