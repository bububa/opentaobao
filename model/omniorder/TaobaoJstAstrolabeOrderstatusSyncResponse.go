package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店派单以及单据相关操作接口 API返回值 
taobao.jst.astrolabe.orderstatus.sync

针对ERP系统部署在门店的商家，将派单状态回流到星盘
*/
type TaobaoJstAstrolabeOrderstatusSyncAPIResponse struct {
    model.CommonResponse
    TaobaoJstAstrolabeOrderstatusSyncResponse
}

// 线下门店派单以及单据相关操作接口 成功返回结果
type TaobaoJstAstrolabeOrderstatusSyncResponse struct {
    XMLName xml.Name `xml:"jst_astrolabe_orderstatus_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // code
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
