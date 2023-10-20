package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsaligenietrackssearch 查询音频
// alibaba.ailabs.aligenie.tracks.search
//
// 搜索类目下的音频信息
func Alibabaailabsaligenietrackssearch(clt *core.SDKClient, req *iot.AlibabaailabsaligenietrackssearchAPIRequest, session string) (*iot.AlibabaailabsaligenietrackssearchAPIResponse, error) {
	var resp iot.AlibabaailabsaligenietrackssearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
