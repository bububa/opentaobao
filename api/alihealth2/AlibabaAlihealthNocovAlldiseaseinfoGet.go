package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthNocovAlldiseaseinfoGet 获取全国疫情统计数据
// alibaba.alihealth.nocov.alldiseaseinfo.get
//
// 获取全国疫情统计数据
func AlibabaAlihealthNocovAlldiseaseinfoGet(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest, resp *alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
