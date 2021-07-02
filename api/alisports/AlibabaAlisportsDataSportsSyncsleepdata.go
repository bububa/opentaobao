package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncsleepdata 阿里体育数据中心用户睡眠数据同步接口
// alibaba.alisports.data.sports.syncsleepdata
//
// 阿里体育数据中心用户睡眠数据同步接口
func AlibabaAlisportsDataSportsSyncsleepdata(clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncsleepdataAPIRequest, session string) (*alisports.AlibabaAlisportsDataSportsSyncsleepdataAPIResponse, error) {
	var resp alisports.AlibabaAlisportsDataSportsSyncsleepdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
