package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// Aliyunindustrytttmplansync 天天特卖生产计划单同步
// aliyun.industry.tttm.plan.sync
//
// ERP系统向天天特卖同步生产计划单的数据
func Aliyunindustrytttmplansync(clt *core.SDKClient, req *tttm.AliyunindustrytttmplansyncAPIRequest, session string) (*tttm.AliyunindustrytttmplansyncAPIResponse, error) {
	var resp tttm.AliyunindustrytttmplansyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
