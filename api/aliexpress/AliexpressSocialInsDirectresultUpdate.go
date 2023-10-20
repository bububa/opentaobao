package aliexpress

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

// Aliexpresssocialinsdirectresultupdate ISV更新INS私信发送的结果
// aliexpress.social.ins.directresult.update
//
// ISV更新INS私信发送的结果
func Aliexpresssocialinsdirectresultupdate(clt *core.SDKClient, req *aliexpress.AliexpresssocialinsdirectresultupdateAPIRequest, session string) (*aliexpress.AliexpresssocialinsdirectresultupdateAPIResponse, error) {
	var resp aliexpress.AliexpresssocialinsdirectresultupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
