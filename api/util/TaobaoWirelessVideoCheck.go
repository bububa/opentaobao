package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoWirelessVideoCheck 无线开放视频内容安全检查
// taobao.wireless.video.check
//
// 无线开放内容检查，提供涉黄暴力政治音视频的异步检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70436.html" target="blank">阿里云内容安全</a>
//
// 此API会进行三个场景的审核，检测不通过的视频将被隐藏，用户无法访问被隐藏的视频。
//
// 目前，该接口仅支持顽兔空间的视频扫描。
func TaobaoWirelessVideoCheck(clt *core.SDKClient, req *util.TaobaoWirelessVideoCheckAPIRequest, session string) (*util.TaobaoWirelessVideoCheckAPIResponse, error) {
	var resp util.TaobaoWirelessVideoCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
