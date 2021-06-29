package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺公司订单信息 API返回值 
taobao.sebp.company.getorderinfo

淘小铺合作公司获取公司订单信息，用于公司结算使用
*/
type TaobaoSebpCompanyGetorderinfoAPIResponse struct {
    model.CommonResponse
    TaobaoSebpCompanyGetorderinfoResponse
}

// 淘小铺公司订单信息 成功返回结果
type TaobaoSebpCompanyGetorderinfoResponse struct {
    XMLName xml.Name `xml:"sebp_company_getorderinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoSebpCompanyGetorderinfoResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
