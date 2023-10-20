package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentsinglevideosubmitauditresult 单视频审核提交审核结果
// yunos.tvpubadmin.content.single.video.submitauditresult
//
// 单视频审核提交审核结果
func Yunostvpubadmincontentsinglevideosubmitauditresult(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentsinglevideosubmitauditresultAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentsinglevideosubmitauditresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
