package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportsdatasportssyncsleepdata 阿里体育数据中心用户睡眠数据同步接口
// alibaba.alisports.data.sports.syncsleepdata
//
// 阿里体育数据中心用户睡眠数据同步接口
func Alibabaalisportsdatasportssyncsleepdata(clt *core.SDKClient, req *alisports.AlibabaalisportsdatasportssyncsleepdataAPIRequest, session string) (*alisports.AlibabaalisportsdatasportssyncsleepdataAPIResponse, error) {
	var resp alisports.AlibabaalisportsdatasportssyncsleepdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
