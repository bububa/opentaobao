package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugdownloadgetentdailytaskdtolist 码上放心数据落地-获取每天日报
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
func Alibabaalihealthdrugdownloadgetentdailytaskdtolist(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
