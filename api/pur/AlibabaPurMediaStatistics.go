package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurmediastatistics 新媒体统计信息
// alibaba.pur.media.statistics
//
// 清博同步新媒体的统计信息给到采购平台
func Alibabapurmediastatistics(clt *core.SDKClient, req *pur.AlibabapurmediastatisticsAPIRequest, session string) (*pur.AlibabapurmediastatisticsAPIResponse, error) {
	var resp pur.AlibabapurmediastatisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
