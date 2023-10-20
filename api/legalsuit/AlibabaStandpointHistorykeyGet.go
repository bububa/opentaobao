package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabastandpointhistorykeyget 查询历史数据
// alibaba.standpoint.historykey.get
//
// 查询历史数据
func Alibabastandpointhistorykeyget(clt *core.SDKClient, req *legalsuit.AlibabastandpointhistorykeygetAPIRequest, session string) (*legalsuit.AlibabastandpointhistorykeygetAPIResponse, error) {
	var resp legalsuit.AlibabastandpointhistorykeygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
