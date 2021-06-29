package dt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫汽车线索产品潜客数据 API返回值 
alibaba.dt.tmllcar.leadsinfo

1.	线索分发是天猫汽车行业流量端最中要的产品，经过前两年的业务和数据端的积累已经对整体业务流程和方案有了清晰的思路；目前数据段已经产沉淀2000W汽车潜客数据，通过运营尝试得到了较好的效果，今年将通过与商家端合作（大搜车-卖车管家）完成潜客分发-商家报价-潜客触达-线索分发-线下核销等一整个汽车人群运营闭环；这个接口反馈大搜车线下门店周围潜客规模及热门车型数据
*/
type AlibabaDtTmllcarLeadsinfoAPIResponse struct {
    model.CommonResponse
    AlibabaDtTmllcarLeadsinfoResponse
}

// 天猫汽车线索产品潜客数据 成功返回结果
type AlibabaDtTmllcarLeadsinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_dt_tmllcar_leadsinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaDtTmllcarLeadsinfoResults `json:"result,omitempty" xml:"result,omitempty"`
}
