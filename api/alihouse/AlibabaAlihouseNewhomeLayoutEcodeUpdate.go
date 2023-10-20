package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomelayoutecodeupdate 新房户型变更E码
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
func Alibabaalihousenewhomelayoutecodeupdate(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomelayoutecodeupdateAPIRequest, session string) (*alihouse.AlibabaalihousenewhomelayoutecodeupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomelayoutecodeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
