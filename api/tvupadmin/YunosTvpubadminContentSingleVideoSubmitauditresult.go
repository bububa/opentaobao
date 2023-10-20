package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentSingleVideoSubmitauditresult 单视频审核提交审核结果
// yunos.tvpubadmin.content.single.video.submitauditresult
//
// 单视频审核提交审核结果
func YunosTvpubadminContentSingleVideoSubmitauditresult(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest, resp *tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
