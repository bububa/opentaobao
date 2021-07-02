package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosStoreGetcloudshelfversion 获取云货架版本信息
// alibaba.mos.store.getcloudshelfversion
//
// 根据屏编号获取云货架版本信息
func AlibabaMosStoreGetcloudshelfversion(clt *core.SDKClient, req *mos.AlibabaMosStoreGetcloudshelfversionAPIRequest, session string) (*mos.AlibabaMosStoreGetcloudshelfversionAPIResponse, error) {
	var resp mos.AlibabaMosStoreGetcloudshelfversionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
