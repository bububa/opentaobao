package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeReviewChangestatus 楼盘测评草稿状态同步
// alibaba.alihouse.newhome.review.changestatus
//
// 楼盘测评草稿状态更新
func AlibabaAlihouseNewhomeReviewChangestatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewChangestatusAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeReviewChangestatusAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeReviewChangestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
