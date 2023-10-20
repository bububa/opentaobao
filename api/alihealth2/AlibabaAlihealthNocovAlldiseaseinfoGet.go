package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthNocovAlldiseaseinfoGet 获取全国疫情统计数据
// alibaba.alihealth.nocov.alldiseaseinfo.get
//
// 获取全国疫情统计数据
func AlibabaAlihealthNocovAlldiseaseinfoGet(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest, session string) (*alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
