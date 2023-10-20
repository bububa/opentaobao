package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseadminthemeupdatestatus 云主题上下架+删除
// alibaba.alihouse.admin.theme.update.status
//
// 云主题上下架+删除
func Alibabaalihouseadminthemeupdatestatus(clt *core.SDKClient, req *alihouse.AlibabaalihouseadminthemeupdatestatusAPIRequest, session string) (*alihouse.AlibabaalihouseadminthemeupdatestatusAPIResponse, error) {
	var resp alihouse.AlibabaalihouseadminthemeupdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
