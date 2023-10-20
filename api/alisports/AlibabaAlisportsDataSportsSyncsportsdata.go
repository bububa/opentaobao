package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportsdatasportssyncsportsdata 阿里体育数据中心用户运动数据同步接口
// alibaba.alisports.data.sports.syncsportsdata
//
// 阿里体育数据中心用户运动数据同步接口
func Alibabaalisportsdatasportssyncsportsdata(clt *core.SDKClient, req *alisports.AlibabaalisportsdatasportssyncsportsdataAPIRequest, session string) (*alisports.AlibabaalisportsdatasportssyncsportsdataAPIResponse, error) {
	var resp alisports.AlibabaalisportsdatasportssyncsportsdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
