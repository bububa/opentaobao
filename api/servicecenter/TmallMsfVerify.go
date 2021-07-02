package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallMsfVerify 喵师傅核销接口
// tmall.msf.verify
//
// msf服务核销的top接口
func TmallMsfVerify(clt *core.SDKClient, req *servicecenter.TmallMsfVerifyAPIRequest, session string) (*servicecenter.TmallMsfVerifyAPIResponse, error) {
	var resp servicecenter.TmallMsfVerifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
