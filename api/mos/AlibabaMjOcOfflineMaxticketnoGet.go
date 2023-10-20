package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocofflinemaxticketnoget pos机获取线下最大小票号
// alibaba.mj.oc.offline.maxticketno.get
//
// 给pos机提供线下最大小票号查询
func Alibabamjocofflinemaxticketnoget(clt *core.SDKClient, req *mos.AlibabamjocofflinemaxticketnogetAPIRequest, session string) (*mos.AlibabamjocofflinemaxticketnogetAPIResponse, error) {
	var resp mos.AlibabamjocofflinemaxticketnogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
