package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqyputpackagebind 码拼箱建立父子关系接口
// alibaba.alihealth.drug.kyt.scqy.putpackagebind
//
// 码拼箱建立父子关系接口
func Alibabaalihealthdrugkytscqyputpackagebind(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqyputpackagebindAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqyputpackagebindAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqyputpackagebindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
