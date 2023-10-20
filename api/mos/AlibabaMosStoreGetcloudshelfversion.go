package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosstoregetcloudshelfversion 获取云货架版本信息
// alibaba.mos.store.getcloudshelfversion
//
// 根据屏编号获取云货架版本信息
func Alibabamosstoregetcloudshelfversion(clt *core.SDKClient, req *mos.AlibabamosstoregetcloudshelfversionAPIRequest, session string) (*mos.AlibabamosstoregetcloudshelfversionAPIResponse, error) {
	var resp mos.AlibabamosstoregetcloudshelfversionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
