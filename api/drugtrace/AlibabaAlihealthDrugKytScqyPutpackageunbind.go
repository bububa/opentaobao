package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqyputpackageunbind 码拼箱解除父子关系接口
// alibaba.alihealth.drug.kyt.scqy.putpackageunbind
//
// 码拼箱解除父子关系接口
func Alibabaalihealthdrugkytscqyputpackageunbind(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqyputpackageunbindAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqyputpackageunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
