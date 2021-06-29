package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
药品激活状态同步 API返回值 
alibaba.alihealth.drug.kyt.queryactivetime

根据赋码资源（CodeVersion + resCode）获得最新激活时间
应用于各地市对接前进行药品目录匹配，医保中心存在的药品可能比较陈旧杂乱
*/
type AlibabaAlihealthDrugKytQueryactivetimeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytQueryactivetimeResponse
}

// 药品激活状态同步 成功返回结果
type AlibabaAlihealthDrugKytQueryactivetimeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_queryactivetime_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihealthDrugKytQueryactivetimeResult `json:"result,omitempty" xml:"result,omitempty"`
}
