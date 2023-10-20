package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportsdatasportssyncuserdata 阿里体育数据中心用户个人信息同步接口
// alibaba.alisports.data.sports.syncuserdata
//
// 阿里体育数据中心用户个人信息同步接口
func Alibabaalisportsdatasportssyncuserdata(clt *core.SDKClient, req *alisports.AlibabaalisportsdatasportssyncuserdataAPIRequest, session string) (*alisports.AlibabaalisportsdatasportssyncuserdataAPIResponse, error) {
	var resp alisports.AlibabaalisportsdatasportssyncuserdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
