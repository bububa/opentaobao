package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsaligeniealbumsget 专辑详情
// alibaba.ailabs.aligenie.albums.get
//
// 给予厂商查询专辑下的音频详情
func Alibabaailabsaligeniealbumsget(clt *core.SDKClient, req *iot.AlibabaailabsaligeniealbumsgetAPIRequest, session string) (*iot.AlibabaailabsaligeniealbumsgetAPIResponse, error) {
	var resp iot.AlibabaailabsaligeniealbumsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
