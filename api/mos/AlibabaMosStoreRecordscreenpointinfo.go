package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosstorerecordscreenpointinfo 云屏埋点数据记录接口
// alibaba.mos.store.recordscreenpointinfo
//
// 记录云屏埋点数据
func Alibabamosstorerecordscreenpointinfo(clt *core.SDKClient, req *mos.AlibabamosstorerecordscreenpointinfoAPIRequest, session string) (*mos.AlibabamosstorerecordscreenpointinfoAPIResponse, error) {
	var resp mos.AlibabamosstorerecordscreenpointinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
