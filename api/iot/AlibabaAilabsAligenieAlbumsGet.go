package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsAligenieAlbumsGet 专辑详情
// alibaba.ailabs.aligenie.albums.get
//
// 给予厂商查询专辑下的音频详情
func AlibabaAilabsAligenieAlbumsGet(clt *core.SDKClient, req *iot.AlibabaAilabsAligenieAlbumsGetAPIRequest, session string) (*iot.AlibabaAilabsAligenieAlbumsGetAPIResponse, error) {
	var resp iot.AlibabaAilabsAligenieAlbumsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
