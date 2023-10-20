package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousebusinessactivitydelete 电商券活动删除
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
func Alibabaalihousebusinessactivitydelete(clt *core.SDKClient, req *alihouse.AlibabaalihousebusinessactivitydeleteAPIRequest, session string) (*alihouse.AlibabaalihousebusinessactivitydeleteAPIResponse, error) {
	var resp alihouse.AlibabaalihousebusinessactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
