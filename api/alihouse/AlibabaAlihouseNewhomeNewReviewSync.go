package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeNewReviewSync 新测评基础信息接口
// alibaba.alihouse.newhome.new.review.sync
//
// 新测评基础信息接口
func AlibabaAlihouseNewhomeNewReviewSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeNewReviewSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeNewReviewSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
