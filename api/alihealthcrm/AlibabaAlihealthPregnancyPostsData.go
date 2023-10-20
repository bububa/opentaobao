package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthpregnancypostsdata 发回帖子信息同步
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
func Alibabaalihealthpregnancypostsdata(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthpregnancypostsdataAPIRequest, session string) (*alihealthcrm.AlibabaalihealthpregnancypostsdataAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthpregnancypostsdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
