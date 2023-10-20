package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncuserdata 阿里体育数据中心用户个人信息同步接口
// alibaba.alisports.data.sports.syncuserdata
//
// 阿里体育数据中心用户个人信息同步接口
func AlibabaAlisportsDataSportsSyncuserdata(clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncuserdataAPIRequest, resp *alisports.AlibabaAlisportsDataSportsSyncuserdataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
