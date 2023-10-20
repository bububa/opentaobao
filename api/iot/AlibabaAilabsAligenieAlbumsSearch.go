package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsaligeniealbumssearch 查询专辑
// alibaba.ailabs.aligenie.albums.search
//
// 搜索类目下的专辑信息
func Alibabaailabsaligeniealbumssearch(clt *core.SDKClient, req *iot.AlibabaailabsaligeniealbumssearchAPIRequest, session string) (*iot.AlibabaailabsaligeniealbumssearchAPIResponse, error) {
	var resp iot.AlibabaailabsaligeniealbumssearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
