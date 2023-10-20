package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqappshieldresultget 用户查询加固结果
// alibaba.security.jaq.app.shieldresult.get
//
// 用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
func Alibabasecurityjaqappshieldresultget(clt *core.SDKClient, req *security.AlibabasecurityjaqappshieldresultgetAPIRequest, session string) (*security.AlibabasecurityjaqappshieldresultgetAPIResponse, error) {
	var resp security.AlibabasecurityjaqappshieldresultgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
