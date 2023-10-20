package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseadminthemeupdate 云主题更新
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
func Alibabaalihouseadminthemeupdate(clt *core.SDKClient, req *alihouse.AlibabaalihouseadminthemeupdateAPIRequest, session string) (*alihouse.AlibabaalihouseadminthemeupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihouseadminthemeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
