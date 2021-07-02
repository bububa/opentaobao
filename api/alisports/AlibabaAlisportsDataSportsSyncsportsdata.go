package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncsportsdata 阿里体育数据中心用户运动数据同步接口
// alibaba.alisports.data.sports.syncsportsdata
//
// 阿里体育数据中心用户运动数据同步接口
func AlibabaAlisportsDataSportsSyncsportsdata(clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncsportsdataAPIRequest, session string) (*alisports.AlibabaAlisportsDataSportsSyncsportsdataAPIResponse, error) {
	var resp alisports.AlibabaAlisportsDataSportsSyncsportsdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
