package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeReviewChangestatus 楼盘测评草稿状态同步
// alibaba.alihouse.newhome.review.changestatus
//
// 楼盘测评草稿状态更新
func AlibabaAlihouseNewhomeReviewChangestatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewChangestatusAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeReviewChangestatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
