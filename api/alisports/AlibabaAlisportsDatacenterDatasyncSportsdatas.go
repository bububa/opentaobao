package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDatacenterDatasyncSportsdatas 阿里体育运动数据同步统一接口
// alibaba.alisports.datacenter.datasync.sportsdatas
//
// 给单方提供同步运动数据到阿里体育的接口
func AlibabaAlisportsDatacenterDatasyncSportsdatas(clt *core.SDKClient, req *alisports.AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest, session string) (*alisports.AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponse, error) {
	var resp alisports.AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
