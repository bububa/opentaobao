package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportsdatasportssyncstatdata 阿里体育数据中心用户当天累积数据同步接口
// alibaba.alisports.data.sports.syncstatdata
//
// 阿里体育数据中心用户当天累积数据同步接口
func Alibabaalisportsdatasportssyncstatdata(clt *core.SDKClient, req *alisports.AlibabaalisportsdatasportssyncstatdataAPIRequest, session string) (*alisports.AlibabaalisportsdatasportssyncstatdataAPIResponse, error) {
	var resp alisports.AlibabaalisportsdatasportssyncstatdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
