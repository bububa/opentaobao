package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthnocovalldiseaseinfoget 获取全国疫情统计数据
// alibaba.alihealth.nocov.alldiseaseinfo.get
//
// 获取全国疫情统计数据
func Alibabaalihealthnocovalldiseaseinfoget(clt *core.SDKClient, req *alihealth2.AlibabaalihealthnocovalldiseaseinfogetAPIRequest, session string) (*alihealth2.AlibabaalihealthnocovalldiseaseinfogetAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthnocovalldiseaseinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
