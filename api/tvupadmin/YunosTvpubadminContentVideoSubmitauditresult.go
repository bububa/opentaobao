package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentvideosubmitauditresult 迎客松提交视频审核结果
// yunos.tvpubadmin.content.video.submitauditresult
//
// 迎客松提交视频审核结果
func Yunostvpubadmincontentvideosubmitauditresult(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentvideosubmitauditresultAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentvideosubmitauditresultAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentvideosubmitauditresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
